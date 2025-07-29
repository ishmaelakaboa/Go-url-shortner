package routes

import (
	"go-url-shortner/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()

	router.POST("/shorten", controllers.ShortenURL)

	router.GET("/:short", controllers.RedirectToOriginal)

	return router
}