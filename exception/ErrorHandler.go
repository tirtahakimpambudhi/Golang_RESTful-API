package exception

import (
	"net/http"
	"restful_api/helper"
	"restful_api/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request , err interface{}) {
	if notFoundError(w,r,err) {
		return
	}

	if validationErrors(w,r,err){
		return
	}

	if entryError(w,r,err) {
		return
	}

	internalServer(w,r,err)
}

func validationErrors(w http.ResponseWriter, r *http.Request , err interface{})bool{
	excep , ok := err.(validator.ValidationErrors)
	if ok {
		w.Header().Set("content-type","application/json")	
		w.WriteHeader(http.StatusBadRequest)
		data := web.ResponseStandar{
			Code: http.StatusBadRequest,
			Message: excep.Error() ,
			Data: nil,
		}
		helper.WriteToResponseBodyError(w,data)
		return true
	} else {
		return false
	}
}
func notFoundError(w http.ResponseWriter, r *http.Request , err interface{})bool{
	excep , ok := err.(ErrorNotFound)
	if ok {
		w.Header().Set("content-type","application/json")
		w.WriteHeader(http.StatusNotFound)
		data := web.ResponseStandar{
			Code: http.StatusNotFound,
			Message: "NOT FOUND",
			Data: excep.Error,
		}
		helper.WriteToResponseBodyError(w,data)
		return true
	} else {
		return false
	}
}
func entryError(w http.ResponseWriter, r *http.Request , err interface{})bool{
	excep , ok := err.(ErrorDuplcated)
	if ok {
		w.Header().Set("content-type","application/json")
		w.WriteHeader(http.StatusConflict)
		data := web.ResponseStandar{
			Code: http.StatusConflict,
			Message: "CONFLICT - RESOURCE ALREADY EXISTS",
			Data: excep.Error,
		}
		helper.WriteToResponseBodyError(w,data)
		return true
	} else {
		return false
	}
}

func internalServer(w http.ResponseWriter, r *http.Request , err interface{}) {
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusInternalServerError)
	data := web.ResponseStandar{
		Code: http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Data: err,
	}
	helper.WriteToResponseBodyError(w,data)
}