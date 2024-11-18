package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	ZipCode     string `json:"cep"`
	PublicPlace string `json:"logradouro"`
	Complement  string `json:"complemento"`
	NeighborHo  string `json:"bairro"`
	Locality    string `json:"localidade"`
	UF          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	http.HandleFunc("/", SearchZipCodeHandler)
	http.ListenAndServe(":8080", nil)

}

func SearchZipCodeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipCodeParam := r.URL.Query().Get("zip")
	if zipCodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	zip, error := SearchZipCode(zipCodeParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// result, erro := json.Marshaler(zip)
	// if error != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// w.Write(result)
	json.NewEncoder(w).Encode(zip)
}

func SearchZipCode(zip string) (*ViaCEP, error) {
	resp, err := http.Get("http://viacep.com.br/ws/" + zip + "/json/")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)

	if error != nil {
		return nil, error
	}
	var c ViaCEP
	error = json.Unmarshal(body, &c)
	if error != nil {
		return nil, error
	}
	return &c, nil

}
