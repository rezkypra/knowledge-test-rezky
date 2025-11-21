package handler

import (
	"ktfs/config"
	"ktfs/model"
	"ktfs/request"
	"ktfs/response"

	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexStudent(c *gin.Context) {
	var students []model.Student
	var data []response.IndexStudent

	err := config.DB.Model(model.Student{}).Find(&students).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	for _, student := range students {
		data = append(data, response.IndexStudent{
			ID: student.ID.String(),
			StudentID: student.StudentID,
			Name: student.Name,
		})
	}

	c.JSON(http.StatusOK, response.Format{
		Code: http.StatusOK,
		Message: "Get students success",
		Data: data,
	})
}

func StoreStudent(c *gin.Context) {
	var input request.StoreStudent
	var student model.Student

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	student = model.Student{
		StudentID: input.StudentID,
		Name:      input.Name,
		Gender:    input.Gender,
		Address:   input.Address,
		EntryYear: input.EntryYear,
		Email:     input.Email,
	}
	err = config.DB.Create(&student).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.Format{
		Code:    http.StatusCreated,
		Message: "Store student success",
	})
}

func ShowStudent(c *gin.Context) {
	var student model.Student
	uuidStudent := c.Param("uuid_student")

	err := config.DB.Where("id = ?", uuidStudent).First(&student).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data := response.ShowStudent{
		ID:        student.ID.String(),
		StudentID: student.StudentID,
		Name:      student.Name,
		Gender:    student.Gender,
		Address:   student.Address,
		EntryYear: student.EntryYear,
	}

	c.JSON(http.StatusOK, response.Format{
		Code:    http.StatusOK,
		Message: "Get student details success",
		Data:    data,
	})
}

func UpdateStudent(c *gin.Context) {
	var student model.Student
	var input request.UpdateStudent
	uuidStudent := c.Param("uuid_student")

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	studentExists, err := model.IsStudentExistsByID(uuidStudent)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	if !studentExists {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: "Student with id " + uuidStudent + " not found",
		})
		return
	}

	err = config.DB.Where("id = ?", uuidStudent).First(&student).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = config.DB.Model(&student).Updates(map[string]interface{}{
		"name":       input.Name,
		"gender":     input.Gender,
		"address":    input.Address,
		"entry_year": input.EntryYear,
		"email":      input.Email,
	}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = config.DB.Where("id = ?", uuidStudent).First(&student).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data := response.ShowStudent{
		ID: student.ID.String(),
		StudentID: student.StudentID,
		Name: student.Name,
		Gender: student.Gender,
		Address: student.Address,
		EntryYear: student.EntryYear,
	}

	c.JSON(http.StatusOK, response.Format{
		Code: http.StatusOK,
		Message: "Update student details success",
		Data: data,
	})
}

func DestroyStudent(c *gin.Context) {
	uuidStudent := c.Param("uuid_student")

	err := config.DB.Where("id = ?", uuidStudent).Delete(&model.Student{}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.Format{
		Code: http.StatusOK,
		Message: "Destroy student data success",
	})
}