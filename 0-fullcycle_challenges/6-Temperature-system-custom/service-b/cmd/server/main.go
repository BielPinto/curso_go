package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/BielPinto/curso_go/0-fullcycle_challenges/5-Temperature-system-by-postal-code/internal/core"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {

	// load .env if present; don't fail the process when it's missing
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", core.SearchCEPHandler)

	http.ListenAndServe(":"+port, r)
	fmt.Printf("Server Ready on %s", port)

}
