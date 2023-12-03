package main

import (
	"fmt"
	"net/http"
	"restful_api/config"
	"restful_api/middleware"
	"restful_api/routes"
)

func main() {
	books := routes.BooksRoutes()
	logging := middleware.Logging{
		Handler: books,
	}
	server := http.Server{
		Addr: config.ADDR,
		Handler: &logging,
	}
	fmt.Printf("Running Server http://%v/api \n",config.ADDR)
	server.ListenAndServe()
}