package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/BielPinto/curso_go/0-fullcycle_challenges/5-Temperature-system-by-postal-code/internal/dto"
)

var ErrCepNotFound = errors.New("cep not found")
var searchViaCep = SearchViaCep
var getTemperature = GetTemperature

func SearchCEPHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")

	if !isValidCEP(cep) {
		fmt.Println("invalid zipcode", http.StatusUnprocessableEntity)
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
	address, err := searchViaCep(cep)
	if err == ErrCepNotFound {
		fmt.Println("can not find zipcode", http.StatusNotFound)
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}
	if err != nil {
		fmt.Println("internal error", http.StatusInternalServerError)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	fmt.Println("adrss:  ", address)
	tempC, err := getTemperature(address.Localidade)
	if err != nil {
		fmt.Println("weather service error", http.StatusInternalServerError)
		http.Error(w, "weather service error", http.StatusInternalServerError)
		return
	}
	response := map[string]float64{
		"temp_C": tempC,
		"temp_F": tempC*1.8 + 32,
		"temp_K": tempC + 273,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func isValidCEP(cep string) bool {
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(cep)
}

func SearchViaCep(cep string) (dto.GetViacepApi, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return dto.GetViacepApi{}, err
	}
	defer resp.Body.Close()

	var data dto.GetViacepApi
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return dto.GetViacepApi{}, err
	}

	if err != nil {
		return dto.GetViacepApi{}, ErrCepNotFound
	}

	return data, nil
}

func GetTemperature(city string) (float64, error) {

	apiKey := os.Getenv("WEATHER_API_KEY")

	fmt.Println("apiKey ", apiKey)
	url := fmt.Sprintf(
		"https://api.weatherapi.com/v1/current.json?key=%s&q=%s",
		apiKey,
		city,
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

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
