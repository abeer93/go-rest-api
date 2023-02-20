package main

import (
	"bytes"
	"fmt"
	"net/http"
	"github.com/abeer93/go-rest-api/handler"
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/abeer93/go-rest-api/book"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"encoding/json"
	"github.com/joho/godotenv"
	bm "github.com/abeer93/go-rest-api/book/models"
)

var bc *book.BookController
var bs *book.BookService
var h *handler.Handler

func TestMain(m *testing.M) {
	fmt.Println(">>>>> main test >>>>>")
	// gin mode test
	gin.SetMode(gin.TestMode)

	// load app env
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// setup DB
	h, err = handler.NewHandler(
		os.Getenv("Test_DB_HOST"),
		os.Getenv("Test_DB_PORT"),
		os.Getenv("Test_DB_NAME"),
		os.Getenv("Test_DB_USER"),
		os.Getenv("Test_DB_PASSWORD"),
	)

	if err != nil {
		panic("Can not connect to database.")
	}

	h.RefreshDatabase()

	bs = book.NewBookService(h)
	bc = book.NewBookController(bs)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func SetUpRouter() *gin.Engine{
    router := gin.Default()
    return router
}

func TestListBooks(t *testing.T) {
	h.SeedDB("book")
	router := SetUpRouter()
	router.GET("/books", bc.ListBooks)

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		fmt.Printf("list books request error >>> %v", err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var books []bm.Book
	json.Unmarshal(res.Body.Bytes(), &books)
	assert.Equal(t, 3, len(books))	
	for _, book := range books {
		assert.NotEmpty(t, book)
	}
}

func TestListBooksWithoutSeed(t *testing.T) {
	router := SetUpRouter()
	router.GET("/books", bc.ListBooks)

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		fmt.Printf("list books request error >>> %v", err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	fmt.Printf("status code %d", res.Code)
	require.Equal(t, http.StatusOK, res.Code)
	var books []bm.Book
	json.Unmarshal(res.Body.Bytes(), &books)
	assert.Equal(t, 0, len(books))
}

func TestCreateBookWithoutSendingData(t *testing.T) {
	router := SetUpRouter()
	router.POST("/books", bc.CreateBook)
	req, err := http.NewRequest("POST", "/books", nil)
	if err != nil {
		fmt.Printf("create book request error >>> %v", err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestCreateBook(t *testing.T) {
	router := SetUpRouter()
	router.POST("/books", bc.CreateBook)
	book := bm.Book{
		Author: "bery",
		Title: "software",
	}
	requestBody, _ := json.Marshal(book)
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("create book request error >>> %v", err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusCreated, res.Code)
}

