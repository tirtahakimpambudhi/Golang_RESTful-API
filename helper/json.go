package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request , result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(result)
	PanicIFError(err)
}



func WriteToResponseBody (w http.ResponseWriter, res interface{}) {
	w.Header().Set("content-type","application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(res)
	PanicIFError(err)
}
func WriteToResponseBodyError (w http.ResponseWriter, res interface{}) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(res)
	PanicIFError(err)
}
