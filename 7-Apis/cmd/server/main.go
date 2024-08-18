package main

import (
	"net/http"

	"github.com/BielPinto/curso_go/7-Apis/configs"
	"github.com/BielPinto/curso_go/7-Apis/infra/database"
	"github.com/BielPinto/curso_go/7-Apis/infra/webserver/handlers"
	"github.com/BielPinto/curso_go/7-Apis/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	config, err := configs.LoadConfig(".")
	println(config.DBDriver)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
