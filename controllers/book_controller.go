package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luannv1203/golang/models"
	"github.com/luannv1203/golang/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetListBooks(c *gin.Context) {
	book := models.Book{}

	listBooks, err := book.FindBooks(c)
	if err != nil {
		responses.JSON(c, http.StatusBadRequest, err, "")
	}

	responses.JSON(c, http.StatusOK, listBooks, "All Books")
}

func CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	newBook := book.Prepare()
	err := book.Validate()
	if err != nil {
		responses.JSON(c, http.StatusBadRequest, nil, "Bad request!")
		return
	}
	res, err := book.CreateBook(newBook)
	if err != nil {
		responses.JSON(c, http.StatusNoContent, nil, "Created Failed")
		return
	}
	responses.JSON(c, http.StatusOK, res, "Create Book Success")
}

func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responses.JSON(c, http.StatusBadRequest, nil, "Bad request!")
		return
	}
	bookRecived, err := book.GetBookByID(objID)
	if err != nil {
		responses.JSON(c, http.StatusOK, nil, "Book not found!")
		return
	}
	responses.JSON(c, http.StatusOK, bookRecived, "")
}
