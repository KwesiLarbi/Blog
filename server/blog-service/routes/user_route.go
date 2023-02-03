package routes

import (
	"github.com/KwesiLarbi/blog-service/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/user/register", controllers.Register())
	router.POST("/user/login", controllers.Login())
}