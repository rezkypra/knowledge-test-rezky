package route

import (
	"ktfs/handler"
	"ktfs/middleware"

	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.Cors())

	publicRoutes := router.Group("/api/v0")
	protectedRoutes := router.Group("/api/v0")
	protectedRoutes.Use(middleware.JwtAuth())

	publicRoutes.POST("/auth/login", handler.Login)
	publicRoutes.POST("/auth/register", handler.Register)
	protectedRoutes.POST("/auth/introspect", handler.Introspect)
	protectedRoutes.POST("/auth/logout", handler.Logout)

	protectedRoutes.GET("/student", handler.IndexStudent)
	protectedRoutes.POST("/student", handler.StoreStudent)
	protectedRoutes.GET("/student/:uuid_student", handler.ShowStudent)
	protectedRoutes.PUT("/student/:uuid_student", handler.UpdateStudent)
	protectedRoutes.DELETE("/student/:uuid_student", handler.DestroyStudent)
	protectedRoutes.POST("/student/subject", handler.AttachSubjectToStudent)

	protectedRoutes.GET("/subject", handler.IndexSubject)
	protectedRoutes.POST("/subject", handler.StoreSubject)
	protectedRoutes.GET("/subject/:uuid_subject", handler.ShowSubject)
	protectedRoutes.PUT("/subject/:uuid_subject", handler.UpdateSubject)
	protectedRoutes.DELETE("/subject/:uuid_subject", handler.DestroySubject)

	return router
}