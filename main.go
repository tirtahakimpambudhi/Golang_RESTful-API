package main

import (
	"fmt"
	"restful_api/config"
	"restful_api/injector"
)

func main() {
	servers := injector.InitializeServers()
	fmt.Printf("Running Server http://%v/api \n", config.ADDR)
	servers.ListenAndServe()
}
