package main

import (
	"encoding/json"
	"fmt"
	"go-book-service/server"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	// Swagger imports
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/aditya-bhuyan/go-book-service/docs"
	_ "github.com/aditya-bhuyan/go-book-service/server" // Import server package for route registration
	// Needed for generated docs
)

// Config holds server configuration properties such as port number
type Config struct {
	ServerPort string `json:"server_port"`
}

// loadConfig reads the server configuration from a JSON file
func loadConfig(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Cannot decode config JSON: %v", err)
	}

	return config
}

// main is the entry point of the server application
// @title           Book Management API
// @version         1.0
// @description     This is a sample server for managing books.
// @host            localhost:8080
// @BasePath        /
func main() {
	// Load configuration from file
	config := loadConfig("config/config.json")

	// Initialize Gin web framework
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register RESTful endpoints
	r.GET("/books", server.GetBooks)          // List all books
	r.POST("/books", server.AddBook)          // Add a new book
	r.DELETE("/books/:id", server.DeleteBook) // Delete a book by ID

	// Start server
	err := r.Run(":" + config.ServerPort)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
