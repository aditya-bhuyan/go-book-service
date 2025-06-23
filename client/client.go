package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	ServerAddress string `json:"server_address"`
}

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func loadConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	return config, err
}

func main() {
	config, err := loadConfig("config/config.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	baseURL := config.ServerAddress

	// Add a book
	book := Book{ID: "1", Title: "Go in Action", Author: "William Kennedy"}
	body, _ := json.Marshal(book)

	resp, err := http.Post(baseURL+"/books", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	fmt.Println("Add Book:", resp.Status)
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	// Get all books
	resp, err = http.Get(baseURL + "/books")
	if err != nil {
		panic(err)
	}
	fmt.Println("\n\nGet Books:")
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
}
