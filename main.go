package main

import (
	"log"
	"os"
	"github.com/abeer93/go-rest-api/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/abeer93/go-rest-api/book"
)

func main() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	h, err := handler.NewHandler(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		panic("Can not connect to database.")
	}

	h.MigrateTable("book")
	// h.SeedDB("book")

	bs := book.NewBookService(h)
	bc := book.NewBookController(bs)
	r := setupRouter(bc)

	r.Run()
}

func setupRouter(bc *book.BookController) *gin.Engine {
	r :=  gin.Default()

	r.GET("/books", bc.ListBooks)
	r.POST("/books", bc.CreateBook)
	r.DELETE("/books/:id", bc.DeleteBook)

	return r
}