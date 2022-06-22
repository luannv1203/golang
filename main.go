package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/luannv1203/golang/config"
	"github.com/luannv1203/golang/routes"
)

var c = config.Config{}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	c.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	// seed.Load(server.DB)
	router := gin.Default()
	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
