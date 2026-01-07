package main

import (
	"log"
	"net/http"

	"github.com/BielPinto/curso_go/0-fullcycle_challenges/5-Temperature-system-by-postal-code/internal/core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", core.SearchCEPHandler)

	http.ListenAndServe(":8080", r)

}
