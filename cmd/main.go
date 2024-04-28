package main

import (
	"log"
	. "github.com/farimarwat/go-books/database"
	. "github.com/farimarwat/go-books/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}
	
	defer Disconnect()
	gin := gin.Default()
	gin.POST("/books",CreateBook)
	gin.GET("/books",ListBooks)
	gin.GET("/books/:name",FindBook)
	gin.Run("localhost:8080")
}

