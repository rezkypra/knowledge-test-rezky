package middleware

import (
	"ktfs/config"
	"ktfs/model"
	"ktfs/response"

	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := ValidateJWT(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Format{
				Code: http.StatusUnauthorized,
				Message: "Authentication required",
			})
			c.Abort()
			return
		}

		user, err := CurrentUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.Format{
				Code: http.StatusUnauthorized,
				Message: "Authentication required",
			})
			c.Abort()
			return
		}

		token := GetTokenFromRequest(c)
		if token != user.Token {
			c.JSON(http.StatusUnauthorized, response.Format{
				Code: http.StatusUnauthorized,
				Message: "Invalid token",
			})
			c.Abort()
			return
		} 

		c.Next()
		return
	}
}

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func ValidateJWT(c *gin.Context) error {
	token, err := GetToken(c)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("The token provided is invalid")
}

func CurrentUser(c *gin.Context) (model.User, error) {
	err := ValidateJWT(c)
	if err != nil {
		return model.User{}, err
	}
	token, _ := GetToken(c)
	claims, _ := token.Claims.(jwt.MapClaims)
	userID, _ := uuid.Parse(claims["id"].(string))

	var user model.User
	err = config.DB.Where("id = ?", userID).Preload("Roles").First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}

func GetToken(c *gin.Context) (*jwt.Token, error) {
	tokenString := GetTokenFromRequest(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func GetTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		if splitToken[0] == "Bearer" {
			return splitToken[1]
		}
	}
	return ""
}