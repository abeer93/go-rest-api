package book

import (
	"github.com/abeer93/go-rest-api/book/models"
	"strconv"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	service Service
}

func NewBookController(service Service) *BookController {
	return &BookController{service}
}

func (bc *BookController) ListBooks(c *gin.Context) {
	books, err := bc.service.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var book book.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	book, err := bc.service.AddNewBook(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	idInt, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	var book book.Book
	book.ID = uint(idInt)
	err := bc.service.RemoveBook(book)
	fmt.Println("response db ", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted Successfully.",
	})
}