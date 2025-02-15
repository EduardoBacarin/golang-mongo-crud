package main

import (
	"fmt"
	"log"
	"os"

	"github.com/EduardoBacarin/golang-mongo-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on .env file load")
	}
	fmt.Println(os.Getenv("APP_NAME") + " is running")

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":" + os.Getenv("APP_PORT")); err != nil {
		log.Fatal(err)
	}
}
