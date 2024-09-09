package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/BielPinto/curso_go/0-fullcycle_challenges/2-Multitheading/internal/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

/*
	MultiThreading an APIs

- Create Function to make request to the APIs.  --> pendent
  - https://brasilapi.com.br/api/cep/v1/01153000 + cep --> pendent
  - http://viacep.com.br/ws/" + cep + "/json/           --> pendent

- Create funciton that saves the fastest reponse and discards the showest one.  --> pendent
- Show the data , address and whitch Api sent it on the command line.  --> pendent
- Limit the response time to 1 second, othewise  the timeout should be displayed. --> pendent
*/

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Get("/", SearchCEPHandler)

	})
	http.HandleFunc("/", SearchCEPHandler)
	http.ListenAndServe(":8080", r)

}

func SearchCEPHandler(w http.ResponseWriter, r *http.Request) {
	channel := make(chan dto.Message)

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
	}

	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go SearchViacep(cepParam, channel)
	go SearchBrasilapi(cepParam, channel)

	select {
	case msg := <-channel:
		fmt.Printf("Received msg: - %s\n", msg)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(msg)
		w.WriteHeader(http.StatusOK)
	case <-time.After(time.Second * 1):
		w.Header().Set("Content-Type", "application/json")
		println("Timeout")
		w.WriteHeader(http.StatusRequestTimeout)
	}

}

// Brasilapi
func SearchBrasilapi(cep string, channel chan<- dto.Message) {
	var cepDto dto.GetBrasilApi
	resp, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)

	if error != nil {
		return
	}
	error = json.Unmarshal(body, &cepDto)
	if error != nil {
		return
	}

	msg := dto.Message{
		Type:         "Brasilapi",
		Cep:          cepDto.Cep,
		State:        cepDto.State,
		City:         cepDto.City,
		Neighborhood: cepDto.Neighborhood,
		Street:       cepDto.Street,
	}
	channel <- msg
}

// Viacep
func SearchViacep(cep string, channel chan<- dto.Message) {
	var cepDto dto.GetViacepApi
	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return
	}
	error = json.Unmarshal(body, &cepDto)
	if error != nil {
		return
	}

	msg := dto.Message{
		Type:         "Viacep",
		Cep:          cepDto.Cep,
		State:        cepDto.Estado,
		City:         cepDto.UF,
		Neighborhood: cepDto.Bairro,
		Street:       cepDto.Logradouro,
	}
	channel <- msg
}
