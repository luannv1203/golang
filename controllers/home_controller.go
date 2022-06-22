package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luannv1203/golang/responses"
)

func Home(c *gin.Context) {
	responses.JSON(c, http.StatusOK, "Welcome to this Awesome API", "")
}
