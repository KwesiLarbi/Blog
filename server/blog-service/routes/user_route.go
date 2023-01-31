package routes

import (
	"github.com/KwesiLarbi/blog-service/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.Register())
}