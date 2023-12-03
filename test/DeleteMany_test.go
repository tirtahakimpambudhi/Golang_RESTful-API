package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"restful_api/helper"
	"testing"
)

func GetManyQueryParams(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Query()

	fmt.Fprintf(w, "isbn :%v", isbn["isbn"] )
}

func TestDeleteMany(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet,"http://localhost:8989?isbn=123&isbn=12344",nil)

	GetManyQueryParams(recorder,request)
	response := recorder.Result()
	body , _:= io.ReadAll(response.Body)
	t.Log(string(body))
}


func TestQueryDeleteMany(t *testing.T){
	isbn := []string{"ISBN1",",ISBN2"}
	query := "DELETE FROM your_table_name WHERE isbn IN (?)"
	query = helper.InQueryPlaceholders(query,6)
	tx := helper.SliceToInterface(isbn)

	t.Log(query)
	t.Log(tx...)
}