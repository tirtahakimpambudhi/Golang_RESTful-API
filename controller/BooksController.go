package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type BooksController interface {
	FindAll(w http.ResponseWriter, r *http.Request , params httprouter.Params) 
	FindByISBN(w http.ResponseWriter, r *http.Request , params httprouter.Params) 
	Create(w http.ResponseWriter, r *http.Request , params httprouter.Params) 
	Creates(w http.ResponseWriter, r *http.Request , params httprouter.Params) 
	Update(w http.ResponseWriter, r *http.Request , params httprouter.Params) 
	Delete(w http.ResponseWriter, r *http.Request , params httprouter.Params) 
	Deletes(w http.ResponseWriter, r *http.Request , params httprouter.Params) 
}