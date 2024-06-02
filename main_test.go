package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetBook(t *testing.T) {
	router := gin.New()
	router.GET("/books/:id", bookController.FindByID)

	req, _ := http.NewRequest("GET", "/books/1", nil)
	req.Header.Set("X-API-KEY", os.Getenv("SECRET_KEY"))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestCreateBook(t *testing.T) {
	router := gin.New()
	router.POST("/books", bookController.Create)

	bookData := `{"title":"Harry Potter","author":"JK Rowling","isbn":"9783503290778","published_date":"1997-06-26"}`
	req, _ := http.NewRequest("POST", "/books", bytes.NewBufferString(bookData))
	req.Header.Set("X-API-KEY", os.Getenv("SECRET_KEY"))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdateBook(t *testing.T) {
	router := gin.New()
	router.PUT("/books/:id", bookController.Update)

	bookData := `{"title":"Harry Potter and the Cursed Child","author":"J. K. Rowling, Jack Thorne, John Tiffany","isbn":"9783503290778","published_date":"2016-07-30"}`
	req, _ := http.NewRequest("PUT", "/books/8", bytes.NewBufferString(bookData))
	req.Header.Set("X-API-KEY", os.Getenv("SECRET_KEY"))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeleteBook(t *testing.T) {
	router := gin.New()
	router.DELETE("/books/:id", bookController.Delete)

	req, _ := http.NewRequest("DELETE", "/books/8", nil)
	req.Header.Set("X-API-KEY", os.Getenv("SECRET_KEY"))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
