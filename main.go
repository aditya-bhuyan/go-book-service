package main

import (
	"encoding/json"
	"fmt"
	"go-book-service/server"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	ServerPort string `json:"server_port"`
}

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

func main() {
	config := loadConfig("config/config.json")
	r := gin.Default()

	r.GET("/books", server.GetBooks)
	r.POST("/books", server.AddBook)
	r.DELETE("/books/:id", server.DeleteBook)

	err := r.Run(":" + config.ServerPort)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
