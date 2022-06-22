package config

import (
	"context"
	"fmt"
	"log"

	"github.com/luannv1203/golang/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	DB *mongo.Database
}

func (c *Config) Initialize(DbUser, DbPassword, DbName string) {
	// credential := options.Credential{
	// 	Username: DbUser,
	// 	Password: DbPassword,
	// }
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017/")

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected MongoDB")
	database := client.Database(DbName)
	models.BookCollection(database)
	return
}
