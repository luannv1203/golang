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
		responses.JSON(c, http.StatusBadRequest, err, "")
	}

	responses.JSON(c, http.StatusOK, listBooks, "All Books")
}
