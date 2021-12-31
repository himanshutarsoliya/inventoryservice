package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/himanshutarsoliya/go-playground/inventoryservice/database"
	"github.com/himanshutarsoliya/go-playground/inventoryservice/product"
	"log"
	"net/http"
)

const basePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(basePath)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
