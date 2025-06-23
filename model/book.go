package model

// Book represents the model for a book resource
// ID is the unique identifier for the book
// Title is the name of the book
// Author is the person who wrote the book
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
