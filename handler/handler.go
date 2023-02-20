package handler

import (
	"github.com/abeer93/go-rest-api/book/models"
	// "strconv"
	"log"
	"gorm.io/driver/postgres"
	"fmt"
	// "net/http"
	// "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// controller --> sevice interface, 
// service -> repo interface
// repo -> model interface

// type Handler interface {
// 	MigrateTable(modelName string)
// 	SeedDB(modelName string)
// 	RefreshDatabase() error 
// }

type Handler struct {
	DB *gorm.DB
}

func NewHandler(dbHost string, dbPort string, dbName string, dbUser string, dbPassword string) (*Handler, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	var h Handler
	db, err := gorm.Open(postgres.Open(connectionString))
	if err == nil {
		// db.AutoMigrate(&book.Book{})
		h.DB = db
	}

	return &h, err
}

func (h *Handler) MigrateTable(modelName string) {
	switch modelName {
		case "book":
			fmt.Print("inside migrate >>>> \n")
			h.DB.AutoMigrate(&book.Book{})

		default:
			break
	}
}

func (h *Handler) SeedDB(modelName string) {
	switch modelName {
		case "book":
			objects := book.GenerateBooksList()
			h.DB.Create(&objects)
		
		default:
			break
	}
}

func (h *Handler) RefreshDatabase() error {
	err := h.DB.Migrator().DropTable(&book.Book{})
	if err != nil {
		return err
	}

	err = h.DB.AutoMigrate(&book.Book{})
	if err != nil {
		return err
	}

	log.Printf("Successfully Refreshed Database.")
	
	return nil
}