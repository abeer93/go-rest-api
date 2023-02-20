package book

import (
	"github.com/abeer93/go-rest-api/book/models"
	"github.com/abeer93/go-rest-api/handler"
	"fmt"
)

type Service interface {
	GetAllBooks() ([]book.Book, error) 
	AddNewBook(book *book.Book) (book.Book, error)
	RemoveBook(book book.Book) error
}

type BookService struct {
	dbHandler *handler.Handler
}

func NewBookService(dbHandler *handler.Handler) *BookService {
	return &BookService{dbHandler}
}

func (bs BookService) GetAllBooks() ([]book.Book, error) {
	var books []book.Book
	err := bs.dbHandler.DB.Find(&books)
	if err.Error != nil {
		return nil, err.Error
	}

	return books, nil
}

func (bs BookService) AddNewBook(book *book.Book) (book.Book, error) {
	result := bs.dbHandler.DB.Create(&book)
	if result.Error != nil {
		return *book, result.Error
	}

	return *book, nil
}

func (bs BookService) RemoveBook(book book.Book) error {
	err := bs.dbHandler.DB.First(&book)
	if err != nil {
		return err.Error
	}
	
	err = bs.dbHandler.DB.Delete(&book)
	fmt.Printf("db error %v ", err.Error)
	if err.Error != nil {
		return err.Error
	}

	return nil
}