package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	DB     *mongo.Database
	Router *mux.Router
}

func (server *Server) Initialize(DbUser, DbPassword, DbName string) {
	credential := options.Credential{
		Username: DbUser,
		Password: DbPassword,
	}
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017/").SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected MongoDB")
	collection := client.Database(DbName)
	server.DB = collection
	server.Router = mux.NewRouter()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 5000")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
