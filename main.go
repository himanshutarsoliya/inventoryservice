package main

import (
	"github.com/himanshutarsoliya/go-playground/inventoryservice/product"
	"log"
	"net/http"
)

const basePath = "/api"

func main() {
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
