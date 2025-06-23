package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	baseURL := "http://localhost:8080"

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
