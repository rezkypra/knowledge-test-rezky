package handler

import (
	"ktfs/config"
	"ktfs/model"
	"ktfs/request"
	"ktfs/response"

	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexSubject(c *gin.Context) {
	var subjects []model.Subject
	var data []response.Subject

	err := config.DB.Model(model.Subject{}).Find(&subjects).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	for _, subject := range subjects {
		data = append(data, response.Subject{
			ID:      subject.ID.String(),
			Name:    subject.Name,
			Credits: subject.Credits,
		})
	}

	c.JSON(http.StatusOK, response.Format{
		Code:    http.StatusOK,
		Message: "Get subjects success",
		Data:    data,
	})
}

func StoreSubject(c *gin.Context) {
	var input request.CreateSubject
	var subject model.Subject

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	subject = model.Subject{
		Name:    input.Name,
		Credits: input.Credits,
	}
	err = config.DB.Create(&subject).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Format{
		Code:    http.StatusCreated,
		Message: "Store subject success",
	})
}

func ShowSubject(c *gin.Context) {
	var subject model.Subject
	uuidSubject := c.Param("uuid_subject")

	err := config.DB.Where("id = ?", uuidSubject).First(&subject).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data := response.Subject{
		ID:      subject.ID.String(),
		Name:    subject.Name,
		Credits: subject.Credits,
	}

	c.JSON(http.StatusOK, response.Format{
		Code:    http.StatusOK,
		Message: "Get subject details success",
		Data:    data,
	})
}

func UpdateSubject(c *gin.Context) {
	var subject model.Subject
	var input request.UpdateSubject
	uuidSubject := c.Param("uuid_subject")

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = config.DB.Where("id = ?", uuidSubject).First(&subject).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = config.DB.Model(&subject).Updates(map[string]interface{}{
		"name":    input.Name,
		"credits": input.Credits,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Format{
		Code:    http.StatusOK,
		Message: "Update subject success",
	})
}

func DestroySubject(c *gin.Context) {
	uuidSubject := c.Param("uuid_subject")

	err := config.DB.Where("id = ?", uuidSubject).Delete(&model.Subject{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Format{
		Code:    http.StatusOK,
		Message: "Destroy subject success",
	})
}

func AttachSubjectToStudent(c *gin.Context) {
	var input request.AttachSubject
	var student model.Student
	var subject model.Subject

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	// Find Student
	err = config.DB.Where("id = ?", input.StudentID).First(&student).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: "Student not found",
		})
		return
	}

	// Find Subject
	err = config.DB.Where("id = ?", input.SubjectID).First(&subject).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: "Subject not found",
		})
		return
	}

	// Attach
	err = config.DB.Model(&student).Association("Subjects").Append(&subject)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Format{
		Code:    http.StatusOK,
		Message: "Subject attached to student success",
	})
}
