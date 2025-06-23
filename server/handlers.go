package server

import (
	"go-book-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// books is an in-memory store for book data
var books = make(map[string]model.Book)

// GetBooks handles GET /books
// Returns a list of all books in the in-memory store
// GetBooks godoc
// @Summary      Get all books
// @Description  Retrieves all books from in-memory storage
// @Tags         books
// @Produce      json
// @Success      200  {array}  model.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	values := make([]model.Book, 0, len(books))
	for _, v := range books {
		values = append(values, v)
	}
	c.JSON(http.StatusOK, values)
}

// AddBook handles POST /books
// Adds a new book to the in-memory store from the request JSON
// AddBook godoc
// @Summary      Add a new book
// @Description  Adds a book to the in-memory store
// @Tags         books
// @Accept       json
// @Produce      json
// @Param        book  body      model.Book  true  "Book to add"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /books [post]
func AddBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books[book.ID] = book
	c.JSON(http.StatusOK, gin.H{"status": "Book added"})
}

// DeleteBook handles DELETE /books/:id
// Removes the book with the given ID from the in-memory store
// DeleteBook godoc
// @Summary      Delete a book
// @Description  Deletes a book by ID from in-memory store
// @Tags         books
// @Param        id   path      string  true  "Book ID"
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	delete(books, id)
	c.JSON(http.StatusOK, gin.H{"status": "Book deleted"})
}
