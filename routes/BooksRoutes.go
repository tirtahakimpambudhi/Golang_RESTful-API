package routes

import (
	"restful_api/controller"
	"restful_api/database"
	"restful_api/exception"
	"restful_api/repository"
	"restful_api/service"


	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func BooksRoutes() *httprouter.Router{
	db := database.GetConnection()
	validaton := validator.New()
	BooksRepository := repository.NewBooksRepository()

	BooksService := service.NewBooksService(BooksRepository,db,validaton)

	BooksController := controller.NewBooksController(BooksService)
	routes := httprouter.New()

	routes.GET("/api/books",BooksController.FindAll)
	routes.POST("/api/books",BooksController.Creates)
	routes.DELETE("/api/books",BooksController.Deletes)

	routes.GET("/api/book/:ISBN",BooksController.FindByISBN)
	routes.POST("/api/book",BooksController.Create)
	routes.PUT("/api/book/:ISBN",BooksController.Update)
	routes.DELETE("/api/book/:ISBN",BooksController.Delete)

	routes.PanicHandler = exception.ErrorHandler
	
	return routes
}

//Ralat Paginaton belum ada
//ISBN Jadi String
//Response nya di samakan dengan api spec
//Logging