package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luannv1203/golang/models"
	"github.com/luannv1203/golang/responses"
)

func GetListBooks(c *gin.Context) {
	book := models.Book{}

	listBooks, err := book.FindBooks(c)
	if err != nil {
		responses.JSON(c.Writer, http.StatusBadRequest, err)
	}

	responses.JSON(c.Writer, http.StatusOK, listBooks)
}
