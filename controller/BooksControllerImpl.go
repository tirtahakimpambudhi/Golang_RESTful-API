package controller

import (
	"log"
	"net/http"
	"restful_api/helper"
	"restful_api/model/web"
	"restful_api/service"

	"github.com/julienschmidt/httprouter"
)

type BooksControllerImpl struct {
	BooksService service.BooksService
}


func (controller *BooksControllerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var book web.RequestBody
	helper.ReadFromRequestBody(r, &book)

	log.Println("Received request to create a book in controller")

	data := controller.BooksService.CreateBook(r.Context(), book)

	response := web.ResponseStandar{
		Code:    200,
		Message: "Successfully created a book",
		Data:    data,
	}

	log.Println("Sending response to client")
	helper.WriteToResponseBody(w, response)
}
func (controller *BooksControllerImpl) Creates(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var books []web.RequestBody
	helper.ReadFromRequestBody(r, &books)

	log.Println("Received request to create a book in controller")

	data := controller.BooksService.CreateBooks(r.Context(), books)

	response := web.ResponseStandar{
		Code:    200,
		Message: "Successfully created a book",
		Data:    data,
	}

	log.Println("Sending response to client")
	helper.WriteToResponseBody(w, response)
}

func (controller *BooksControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	page := r.URL.Query().Get("page")

	log.Printf("Received request to find all books in controller with page: %v\n", page)

	pagination := controller.BooksService.Pagination(r.Context(), page)
	books := controller.BooksService.FindAllBooks(r.Context(), pagination)
	data := web.ResponseArray{
		Books: books,
		Page:  pagination,
	}
	response := web.ResponseStandar{
		Code:    200,
		Message: "Successfully found all books",
		Data:    data,
	}

	log.Println("Sending response to client")
	helper.WriteToResponseBody(w, response)
}

func (controller *BooksControllerImpl) FindByISBN(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ISBN := params.ByName("ISBN")

	log.Printf("Received request to find a book by ISBN (%v) in controller\n", ISBN)

	data := controller.BooksService.FindBookByISBN(r.Context(), ISBN)
	response := web.ResponseStandar{
		Code:    200,
		Message: "Successfully found a book by ISBN",
		Data:    data,
	}

	log.Println("Sending response to client")
	helper.WriteToResponseBody(w, response)
}

func (controller *BooksControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var book web.RequestBody
	helper.ReadFromRequestBody(r, &book)
	ISBN := params.ByName("ISBN")

	log.Printf("Received request to update a book with ISBN (%v) in controller\n", ISBN)

	data := controller.BooksService.UpdateBook(r.Context(), ISBN, book)
	response := web.ResponseStandar{
		Code:    200,
		Message: "Successfully updated a book",
		Data:    data,
	}

	log.Println("Sending response to client")
	helper.WriteToResponseBody(w, response)
}

func (controller *BooksControllerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ISBN := params.ByName("ISBN")

	log.Printf("Received request to delete a book with ISBN (%v) in controller\n", ISBN)

	controller.BooksService.DeleteBook(r.Context(), ISBN)
	response := web.ResponseStandar{
		Code:    200,
		Message: "Successfully deleted a book",
		Data:    nil,
	}

	log.Println("Sending response to client")
	helper.WriteToResponseBody(w, response)
}
func (controller *BooksControllerImpl) Deletes(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ISBNs := r.URL.Query()["ISBN"]

	log.Printf("Received request to delete a book with ISBNs (%v) in controller\n", helper.SliceToInterface(ISBNs)...)

	controller.BooksService.DeleteBooks(r.Context(),ISBNs)
	response := web.ResponseStandar{
		Code:    200,
		Message: "Successfully deleted a books",
		Data:    nil,
	}

	log.Println("Sending response to client")
	helper.WriteToResponseBody(w, response)
}


func NewBooksController( service service.BooksService)BooksController{
	return &BooksControllerImpl{
		BooksService: service,
	}
}