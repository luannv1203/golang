package models

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
}

var collection *mongo.Collection

func BookCollection(c *mongo.Database) {
	collection = c.Collection("books")
}

func (b *Book) FindBooks(c *gin.Context) (*[]Book, error) {
	c.DefaultQuery("page", "0")
	c.DefaultQuery("size", "10")
	fmt.Println(c.DefaultQuery("page", "0"))
	fmt.Println(c.DefaultQuery("size", "10"))
	books := []Book{}
	list, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		return &books, err
	}

	for list.Next(context.TODO()) {
		var book Book
		list.Decode(&book)
		books = append(books, book)
	}

	return &books, err
}
