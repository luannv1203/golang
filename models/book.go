package models

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty"`
	Author string             `json:"author,omitempty"`
}

var collection *mongo.Collection

func BookCollection(c *mongo.Database) {
	collection = c.Collection("books")
}

func (b *Book) Prepare() Book {
	return Book{
		Id:     primitive.NewObjectID(),
		Title:  b.Title,
		Author: b.Author,
	}
}

func (b *Book) Validate() error {
	if b.Title == "" {
		return errors.New("Required Title!")
	}
	if b.Author == "" {
		return errors.New("Required Author!")
	}
	return nil
}

func (b *Book) FindBooks(c *gin.Context) (*[]Book, error) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "1"))
	fmt.Println(page, size)
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

func (b *Book) CreateBook(newBook Book) (*Book, error) {
	_, err := collection.InsertOne(context.TODO(), newBook)

	if err != nil {
		return &Book{}, err
	}

	return &newBook, nil
}

func (b *Book) GetBookByID(id primitive.ObjectID) (*Book, error) {
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&b)

	if err != nil {
		return nil, err
	}

	return b, nil
}
