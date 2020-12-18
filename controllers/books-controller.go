package controllers

import (
	"github.com/apichef/blink-restful-api-go/dtos"
	"github.com/apichef/blink-restful-api-go/models"
	"github.com/apichef/blink-restful-api-go/repositories"
	"github.com/apichef/blink-restful-api-go/responses"
	"github.com/gin-gonic/gin"
)

type BooksController interface {
	GetAll(context *gin.Context)
	GetBook(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type booksController struct {
	repository repositories.BooksRepository
}

func New(repository repositories.BooksRepository) BooksController {
	return &booksController {
		repository: repository,
	}
}

// GetAll godoc
// @Summary Get books
// @Description List all available books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} dtos.Book
// @Failure 500 {object} httputil.HTTPError
// @Router /books [get]
func (controller *booksController) GetAll(context *gin.Context) {
	responses.Ok(context, controller.repository.GetAll())
}

// GetBook godoc
// @Summary Show a book
// @Description get book by ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param book path int true "Book ID"
// @Success 200 {object} dtos.Book
// @Failure 500 {object} httputil.HTTPError
// @Router /books/{book} [get]
func (controller *booksController) GetBook(context *gin.Context) {
	responses.Ok(context, controller.repository.GetBook(context.Param("book")))
}

// Create godoc
// @Summary Add a book
// @Description add by json book
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body dtos.CreateBookDto true "Add book"
// @Success 200 {object} dtos.Book
// @Failure 422 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /books [post]
func (controller *booksController) Create(context *gin.Context) {
	var createBook dtos.CreateBookDto

	if err := context.ShouldBindJSON(&createBook); err != nil {
		responses.BadRequest(context, err)
		return
	}

	book := models.Book {
		Name: createBook.Name,
		Description: createBook.Description,
		GenreID: createBook.GenreID,
		PublisherID: createBook.PublisherID,
	}

	responses.Created(context, controller.repository.Create(book))
}

func (controller *booksController) Update(context *gin.Context) {
	var updateBook dtos.UpdateBookDto
	err := context.ShouldBindJSON(&updateBook)

	if err!= nil {
		responses.BadRequest(context, err)
		return
	}

	book := controller.repository.GetBook(context.Param("book"))

	book.Name = updateBook.Name
	book.Description = updateBook.Description
	book.GenreID = updateBook.GenreID
	book.PublisherID = updateBook.PublisherID

	controller.repository.Update(book)

	responses.NoContent(context)
}

func (controller *booksController) Delete(context *gin.Context) {
	book := controller.repository.GetBook(context.Param("book"))

	controller.repository.Delete(book)

	responses.NoContent(context)
}