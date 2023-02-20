package book

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GenerateBooksList() []Book {
	return []Book{
		{Title: "Harry Potter", Author: "J. K. Rowling"},
		{Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
		{Title: "The Wizard of Oz", Author: "L. Frank Baum"},
	}
}
