package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Config holds server configuration details
type Config struct {
	ServerAddress string `json:"server_address"`
}

// Book represents a book resource
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// loadConfig reads the JSON configuration from file
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

// prompt displays available actions and returns user choice
func prompt() string {
	fmt.Println("\nChoose an action:")
	fmt.Println("1. Add Book")
	fmt.Println("2. Get All Books")
	fmt.Println("3. Delete Book")
	fmt.Println("4. Exit")
	fmt.Print("Enter choice [1-4]: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// readBookDetails prompts user to input book fields
func readBookDetails() Book {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Book ID: ")
	scanner.Scan()
	id := scanner.Text()

	fmt.Print("Enter Title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter Author: ")
	scanner.Scan()
	author := scanner.Text()

	return Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
}

// addBook sends a POST request to add a new book
func addBook(baseURL string, book Book) {
	body, _ := json.Marshal(book)

	resp, err := http.Post(baseURL+"/books", "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error adding book:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("Add Book:", resp.Status)
	io.Copy(os.Stdout, resp.Body)
}

// getBooks sends a GET request to retrieve all books
func getBooks(baseURL string) {
	resp, err := http.Get(baseURL + "/books")
	if err != nil {
		fmt.Println("Error fetching books:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("\nBook List:")
	io.Copy(os.Stdout, resp.Body)
}

// deleteBook sends a DELETE request using the book ID
func deleteBook(baseURL string, id string) {
	req, err := http.NewRequest(http.MethodDelete, baseURL+"/books/"+id, nil)
	if err != nil {
		fmt.Println("Error creating delete request:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error deleting book:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Delete Book:", resp.Status)
	io.Copy(os.Stdout, resp.Body)
}

func main() {
	config, err := loadConfig("config/config.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	baseURL := config.ServerAddress
	scanner := bufio.NewScanner(os.Stdin)

	for {
		choice := prompt()
		switch choice {
		case "1":
			book := readBookDetails()
			addBook(baseURL, book)
		case "2":
			getBooks(baseURL)
		case "3":
			fmt.Print("Enter Book ID to delete: ")
			scanner.Scan()
			deleteBook(baseURL, scanner.Text())
		case "4":
			fmt.Println("Exiting client.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
