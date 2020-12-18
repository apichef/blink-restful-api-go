package repositories

import (
	dbConnection "github.com/apichef/blink-restful-api-go/database"
	"github.com/apichef/blink-restful-api-go/models"
	"gorm.io/gorm"
)

type BooksRepository interface {
	GetAll() []models.Book
	GetBook(ID string) models.Book
	Create(book models.Book) models.Book
	Update(book models.Book)
	Delete(book models.Book)
}

type database struct {
	connection *gorm.DB
}

func New() BooksRepository {
	return &database {
		connection: dbConnection.Open(),
	}
}

func (db *database) GetAll() []models.Book {
	var books []models.Book
	db.connection.Find(&books)

	return books
}

func (db *database) GetBook(ID string) models.Book {
	var book models.Book

	db.connection.First(&book, ID)

	return book
}

func (db *database) Create(book models.Book) models.Book {
	db.connection.Create(&book)

	return book
}

func (db *database) Update(book models.Book)  { 
	db.connection.Save(&book)
}

func (db *database) Delete(book models.Book)  {
	db.connection.Delete(&book)
}