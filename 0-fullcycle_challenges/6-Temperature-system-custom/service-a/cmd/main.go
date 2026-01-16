package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/BielPinto/curso_go/0-fullcycle_challenges/6-Temperature-system-custom/service-a/internal/telemetry"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type CEPRequest struct {
	CEP string `json:"cep"`
}

type TemperatureResponse struct {
	City  string  `json:"city"`
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

var tracer = otel.Tracer("service-a")

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize OTEL tracer
	tp, err := telemetry.InitTracerProvider(ctx, "service-a")
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
		port = "8081"
	}

	serviceBURL := os.Getenv("SERVICE_B_URL")
	if serviceBURL == "" {
		serviceBURL = "http://service-b:8080"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Endpoint that receives CEP via POST and calls Service B
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		requestCtx, span := tracer.Start(r.Context(), "POST /")
		defer span.End()

		var req CEPRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			span.AddEvent("failed to decode request")
			span.SetStatus(codes.Error, "decode error")
			fmt.Println("Error decoding request:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid zipcode"})
			return
		}

		fmt.Printf("Received CEP request: %s\n", req.CEP)

		// Validate CEP: must be 8 digits and a string
		if !isValidCEP(req.CEP) {
			span.AddEvent("invalid CEP format")
			span.SetStatus(codes.Error, "invalid CEP")
			fmt.Println("Invalid CEP format:", req.CEP)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid zipcode"})
			return
		}

		// Create span for Service B call
		callCtx, callSpan := tracer.Start(requestCtx, "call_service_b")
		defer callSpan.End()

		// Call Service B to get temperature
		reqWithCtx, _ := http.NewRequestWithContext(callCtx, "GET", serviceBURL+"/?cep="+req.CEP, nil)
		resp, err := http.DefaultClient.Do(reqWithCtx)
		if err != nil {
			callSpan.AddEvent("error calling service B")
			callSpan.SetStatus(codes.Error, err.Error())
			fmt.Println("Error calling Service B:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal error"})
			return
		}
		defer resp.Body.Close()

		callSpan.AddEvent("received response from service B",
			trace.WithAttributes(attribute.Int("status_code", resp.StatusCode)),
		)

		// If Service B returns 422 or 404, pass it through
		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(resp.StatusCode)
			w.Write(body)
			return
		}

		// Read response from Service B
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			span.AddEvent("error reading response")
			span.SetStatus(codes.Error, "read error")
			fmt.Println("Error reading response:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal error"})
			return
		}

		// Parse and forward the response
		var tempData map[string]interface{}
		if err := json.Unmarshal(body, &tempData); err != nil {
			span.AddEvent("error parsing response")
			span.SetStatus(codes.Error, "parse error")
			fmt.Println("Error parsing response:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal error"})
			return
		}

		span.AddEvent("successfully processed request")

		// Return the temperature data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tempData)
	})

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		healthCtx, span := tracer.Start(r.Context(), "GET /health")
		defer span.End()
		_ = healthCtx

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	fmt.Printf("Service A Ready on port %s\n", port)

	// Create a channel to handle graceful shutdown
	done := make(chan struct{})
	go func() {
		http.ListenAndServe(":"+port, r)
	}()

	// Wait for signal
	<-done
}

// Validate CEP: must be exactly 8 digits and a string
func isValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}
