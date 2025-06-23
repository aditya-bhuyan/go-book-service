package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aditya-bhuyan/go-book-service/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupRouter initializes routes for testing
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/books", GetBooks)
	router.POST("/books", AddBook)
	router.DELETE("/books/:id", DeleteBook)
	return router
}

// TestAddBook validates POST /books
func TestAddBook(t *testing.T) {
	router := setupRouter()

	book := model.Book{ID: "101", Title: "Clean Code", Author: "Robert C. Martin"}
	body, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

// TestGetBooks validates GET /books
func TestGetBooks(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/books", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

// TestDeleteBook validates DELETE /books/:id
func TestDeleteBook(t *testing.T) {
	router := setupRouter()

	// First add a book to delete
	book := model.Book{ID: "202", Title: "The Go Programming Language", Author: "Alan Donovan"}
	body, _ := json.Marshal(book)
	addReq, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	addReq.Header.Set("Content-Type", "application/json")
	addResp := httptest.NewRecorder()
	router.ServeHTTP(addResp, addReq)

	// Now delete it
	req, _ := http.NewRequest("DELETE", "/books/202", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
