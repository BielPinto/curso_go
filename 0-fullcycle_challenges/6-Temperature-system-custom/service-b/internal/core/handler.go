package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"

	"github.com/BielPinto/curso_go/0-fullcycle_challenges/5-Temperature-system-by-postal-code/internal/dto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var ErrCepNotFound = errors.New("cep not found")
var searchViaCep = SearchViaCep
var getTemperature = GetTemperature
var tracer = otel.Tracer("service-b")

func SearchCEPHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "GET /")
	defer span.End()

	cep := r.URL.Query().Get("cep")
	span.AddEvent("received CEP",
		trace.WithAttributes(attribute.String("cep", cep)),
	)

	if !isValidCEP(cep) {
		span.AddEvent("invalid CEP format")
		span.SetStatus(codes.Error, "invalid CEP")
		fmt.Println("invalid zipcode", http.StatusUnprocessableEntity)
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	// Create span for ViaCEP search
	searchCtx, searchSpan := tracer.Start(ctx, "search_viacep")
	address, err := searchViaCep(searchCtx, cep)
	searchSpan.End()

	if err == ErrCepNotFound {
		span.AddEvent("CEP not found")
		span.SetStatus(codes.Error, "CEP not found")
		fmt.Println("can not find zipcode", err)
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}
	if err != nil {
		span.AddEvent("error searching CEP")
		span.SetStatus(codes.Error, err.Error())
		fmt.Println("internal error", http.StatusInternalServerError)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	fmt.Println("adrss:  ", address)

	// Create span for temperature lookup
	tempCtx, tempSpan := tracer.Start(ctx, "get_temperature")
	tempC, err := getTemperature(tempCtx, address.Localidade)
	tempSpan.End()

	if err != nil {
		span.AddEvent("error fetching temperature")
		span.SetStatus(codes.Error, err.Error())
		fmt.Println("weather service error", http.StatusInternalServerError)
		http.Error(w, "weather service error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"city":   address.Localidade,
		"temp_C": tempC,
		"temp_F": tempC*1.8 + 32,
		"temp_K": tempC + 273,
	}

	span.AddEvent("successfully retrieved temperature",
		trace.WithAttributes(
			attribute.String("city", address.Localidade),
			attribute.Float64("temp_C", tempC),
		),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func isValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

func SearchViaCep(ctx context.Context, cep string) (dto.GetViacepApi, error) {
	req, _ := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/"+cep+"/json/", nil)
	resp, err := http.DefaultClient.Do(req)
	fmt.Println("viacep   -> cep: ", cep)
	fmt.Println("resp from viacep: ", resp.Body)
	if err != nil {
		return dto.GetViacepApi{}, err
	}
	defer resp.Body.Close()

	var data dto.GetViacepApi
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return dto.GetViacepApi{}, err
	}

	if data.Cep == "" {
		return dto.GetViacepApi{}, ErrCepNotFound
	}

	return data, nil
}

func GetTemperature(ctx context.Context, city string) (float64, error) {

	apiKey := os.Getenv("WEATHER_API_KEY")
	fmt.Println("WEATHER_API_KEY: ", apiKey)
	if apiKey == "" {
		fmt.Println("WEATHER_API_KEY not set")
		return 0, fmt.Errorf("WEATHER_API_KEY not set")
	}

	// URL encode the city name to handle special characters like "ã"
	encodedCity := url.QueryEscape(city)

	urlStr := fmt.Sprintf(
		"https://api.weatherapi.com/v1/current.json?key=%s&q=%s",
		apiKey,
		encodedCity,
	)
	fmt.Println("url: ", urlStr)

	req, _ := http.NewRequestWithContext(ctx, "GET", urlStr, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("GetTemperature - error", err)
		return 0, err
	}
	defer resp.Body.Close()
	fmt.Println("resp.StatusCode", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		fmt.Println("weather api error:", string(body))
		return 0, fmt.Errorf("weather api error: %s", body)
	}

	var result struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	fmt.Println("resp.Body", result.Current)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("error", err)
		return 0, err
	}

	return result.Current.TempC, nil
}
