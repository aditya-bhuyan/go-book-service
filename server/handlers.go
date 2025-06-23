package server

import (
	"go-book-service/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var books = make(map[string]model.Book)

func GetBooks(c *gin.Context) {
	values := make([]model.Book, 0, len(books))
	for _, v := range books {
		values = append(values, v)
	}
	c.JSON(http.StatusOK, values)
}

func AddBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books[book.ID] = book
	c.JSON(http.StatusOK, gin.H{"status": "Book added"})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	delete(books, id)
	c.JSON(http.StatusOK, gin.H{"status": "Book deleted"})
}
