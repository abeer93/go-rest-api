package book

import (
	"github.com/abeer93/go-rest-api/book/models"
	"github.com/abeer93/go-rest-api/handler"
)

type Service interface {
	GetAllBooks() ([]book.Book, error) 
	AddNewBook(bk *book.Book) (book.Book, error)
	RemoveBook(ID int64) error
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

func (bs BookService) AddNewBook(bk *book.Book) (book.Book, error) {
	result := bs.dbHandler.DB.Create(&bk)
	if result.Error != nil {
		return *bk, result.Error
	}

	return *bk, nil
}

func (bs BookService) RemoveBook(ID int64) error {
	var bk book.Book
	// err := bs.dbHandler.DB.First(&b)
	// if err != nil {
	// 	return err.Error
	// }

	err := bs.dbHandler.DB.Where("id = ?", ID).Delete(&bk)
	if err.Error != nil {
		return err.Error
	}

	return nil
}