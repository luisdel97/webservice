package main

import (
	"net/http"
	"webservice/product"
)

const apiBasePath = "/api"

func main() {

	product.SetupRoutes(apiBasePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		return 
	}
}
