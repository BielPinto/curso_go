package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/BielPinto/curso_go/0-fullcycle_challenges/5-Temperature-system-by-postal-code/internal/core"
	"github.com/BielPinto/curso_go/0-fullcycle_challenges/5-Temperature-system-by-postal-code/internal/telemetry"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// load .env if present; don't fail the process when it's missing
	godotenv.Load()

	// Initialize OTEL tracer
	tp, err := telemetry.InitTracerProvider(ctx, "service-b")
	if err != nil {
		fmt.Printf("Failed to initialize tracer: %v\n", err)
	}
	defer func() {
		if err := telemetry.Shutdown(ctx, tp); err != nil {
			fmt.Printf("Failed to shutdown tracer: %v\n", err)
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", core.SearchCEPHandler)

	fmt.Printf("Server Ready on %s\n", port)
	http.ListenAndServe(":"+port, r)
}
