package handler

import (
	"ktfs/config"
	"ktfs/model"
	"ktfs/middleware"
	"ktfs/request"
	"ktfs/response"

	"errors"
	"net/http"
	"os"
	"strconv"
	"time"
	
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	jwt "github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var input request.Login
	var user model.User

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userExists, err := model.IsUserExistsByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if userExists {
		err = config.DB.Where("email = ?", input.Email).First(&user).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Format{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		err = user.ValidatePassword(input.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.Format{
				Code: http.StatusBadRequest,
				Message: "Invalid user credentials",
			})
			return
		}

	} else {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: "Invalid user credentials",
		})
		return
	}

	jwt, err := GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = config.DB.Model(&user).Updates(map[string]interface{}{"token": jwt,}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	jwtTTL, _ := strconv.Atoi(os.Getenv("JWT_TTL"))
	data := response.Login{
		Token: user.Token,
		TokenType: "Bearer",
		ExpiresIn: jwtTTL,
	}

	c.JSON(http.StatusOK, response.Format{
		Code: http.StatusOK,
		Message: "Login success",
		Data: data,
	})
}

func Register(c *gin.Context) {
	var input request.Register
	var user model.User

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	user = model.User{
		Name: input.Name,
		Email: input.Email,
		Password:	GeneratePassword(input.Password),
	}

	err = config.DB.Create(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, response.Format{
				Code: http.StatusBadRequest,
				Message: "Email already taken",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, response.Format{
				Code: http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}
	}
	
	c.JSON(http.StatusOK, response.Format{
		Code: http.StatusOK,
		Message: user.Email+" registered successfully, please login using the registered credentials",
	})
}

func Introspect(c *gin.Context) {
	token, _ := middleware.GetToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	user, err := middleware.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	data := response.TokenIntrospect{
		ExpireAt: time.Unix(int64(claims["eat"].(float64)), 0),
		ID: user.ID.String(),
		Name: user.Name,
		Email: user.Email,
	}

	c.JSON(http.StatusOK, response.Format{
		Code: http.StatusOK,
		Message: "Introspect success",
		Data: data,
	})
}

func Logout(c *gin.Context) {
	user, err := middleware.CurrentUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = config.DB.Model(&user).Updates(map[string]interface{}{"token": nil,}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Format{
			Code: http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, response.Format{
		Code: http.StatusOK,
		Message: user.Name+" logged out",
	})
}

func GeneratePassword(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password
	}
	return string(passwordHash)
}

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user model.User) (string, error) {
	jwtTTL, _ := strconv.Atoi(os.Getenv("JWT_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(jwtTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}