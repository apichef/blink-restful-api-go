package routes

import (
	"github.com/apichef/blink-restful-api-go/controllers"
	"github.com/apichef/blink-restful-api-go/repositories"
	"github.com/gin-gonic/gin"
)

var(
	boobsRepository repositories.BooksRepository = repositories.New()
	booksController controllers.BooksController = controllers.New(boobsRepository)
)

func AddBookRoutes(rg *gin.RouterGroup) {
	books := rg.Group("/books")
	books.GET("/", booksController.GetAll)
	books.GET("/:book", booksController.GetBook)
	books.POST("/", booksController.Create)
	books.PUT("/:book", booksController.Update)
	books.DELETE("/:book", booksController.Delete)
}
