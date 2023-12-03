package middleware

import (
	"log"
	"net/http"
)

type Logging struct {
	Handler 	http.Handler
}

func (lg *Logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v",r.Method,r.URL.Path)
	lg.Handler.ServeHTTP(w,r)
}