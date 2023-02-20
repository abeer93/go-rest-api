package book

import (
	"bytes"
	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/mock"
	"math/rand"
	"encoding/json"
	book "github.com/abeer93/go-rest-api/book/models"
	"github.com/abeer93/go-rest-api/handler"
	mocks "github.com/abeer93/go-rest-api/mocks/book"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println(">>>>> main test >>>>>")

	// gin mode test
	gin.SetMode(gin.TestMode)

	// load app env
	err := godotenv.Load("../app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	// setup DB
	h, err := handler.NewHandler(
		os.Getenv("Test_DB_HOST"),
		os.Getenv("Test_DB_PORT"),
		os.Getenv("Test_DB_NAME"),
		os.Getenv("Test_DB_USER"),
		os.Getenv("Test_DB_PASSWORD"),
	)
	if err != nil {
		panic("Can not connect to database.")
	}

	// refresh the database (drop all tables and migrate it again)
	h.RefreshDatabase()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// instantiate book service mocker and book controller
func instantiateController(t *testing.T) (*mocks.Service, *BookController) {
	bsMock := mocks.NewService(t)
	bc := NewBookController(bsMock)

	return bsMock, bc
}

func TestListBooks(t *testing.T) {
	bsMock, bc := instantiateController(t)

	// setup router
	router := SetUpRouter()
	router.GET("/books", bc.ListBooks)

	books := book.GenerateBooksList()
	bsMock.On("GetAllBooks").Return(books, nil)

	// prepare request
	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		fmt.Printf("list books request error >>> %v", err)
	}

	// send request
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	// assertions
	assert.Equal(t, http.StatusOK, res.Code)

	var booksList []book.Book
	json.Unmarshal(res.Body.Bytes(), &booksList)
	assert.Equal(t, 3, len(booksList))
	for _, book := range booksList {
		assert.NotEmpty(t, book)
	}
}

func TestListBooksWithoutDataExist(t *testing.T) {
	bsMock, bc := instantiateController(t)

	router := SetUpRouter()
	router.GET("/books", bc.ListBooks)

	var books []book.Book
	bsMock.On("GetAllBooks").Return(books, nil)

	req, err := http.NewRequest("GET", "/books", nil)
	if err != nil {
		fmt.Printf("list books request error >>> %v", err)
	}
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	require.Equal(t, http.StatusOK, res.Code)
	var booksList []book.Book
	json.Unmarshal(res.Body.Bytes(), &booksList)
	assert.Equal(t, 0, len(booksList))
}

func TestCreateBook(t *testing.T) {
	bsMock, bc := instantiateController(t)

	bk := book.Book{
		Author: "bery",
		Title:  "software",
	}
	bsMock.On("AddNewBook", &bk).Return(bk, nil)

	router := SetUpRouter()
	router.POST("/books", bc.CreateBook)
	requestBody, _ := json.Marshal(bk)
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Printf("create book request error >>> %v", err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	var createdBook book.Book
	assert.Equal(t, http.StatusCreated, res.Code)
	json.Unmarshal(res.Body.Bytes(), &createdBook)
	assert.Equal(t, bk, createdBook)
}

func TestDeleteBook(t *testing.T) {
	bsMock, bc := instantiateController(t)

	randomNumber := rand.Intn(10)
	id := int64(randomNumber)
	bsMock.On("RemoveBook", id).Return(nil)

	router := SetUpRouter()
	router.DELETE("/books/:id", bc.DeleteBook)
	url := fmt.Sprint("/books/", id)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Printf("remove book request error >>> %v", err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
}
