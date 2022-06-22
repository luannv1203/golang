package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luannv1203/golang/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", controllers.Home)
	router.GET("/books", controllers.GetListBooks)
}
