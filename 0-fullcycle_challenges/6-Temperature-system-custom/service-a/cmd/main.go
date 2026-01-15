package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

func main() {
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
		var req CEPRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println("Error decoding request:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid zipcode"})
			return
		}

		// Validate CEP: must be 8 digits and a string
		if !isValidCEP(req.CEP) {
			fmt.Println("Invalid CEP format:", req.CEP)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(map[string]string{"message": "invalid zipcode"})
			return
		}

		// Call Service B to get temperature
		resp, err := http.Get(serviceBURL + "/?cep=" + req.CEP)
		if err != nil {
			fmt.Println("Error calling Service B:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal error"})
			return
		}
		defer resp.Body.Close()

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
			fmt.Println("Error reading response:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal error"})
			return
		}

		// Parse and forward the response
		var tempData map[string]interface{}
		if err := json.Unmarshal(body, &tempData); err != nil {
			fmt.Println("Error parsing response:", err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"message": "internal error"})
			return
		}

		// Return the temperature data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(tempData)
	})

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	fmt.Printf("Service A Ready on port %s\n", port)
	http.ListenAndServe(":"+port, r)
}

// Validate CEP: must be exactly 8 digits and a string
func isValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}
