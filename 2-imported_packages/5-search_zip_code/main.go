package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

	for _, zip := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + zip + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error  while making request%v \n", err)
		}

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while reading response: %v \n", err)
		}
		var data ViaCEP
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while pasting the response: %v \n", err)
		}

		fmt.Println(data)
		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while creating file %v \n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("Zip: %s locality: %s, UF: %s", data.ZipCode, data.Locality, data.UF))

		fmt.Println("File created successfully!")
		fmt.Println("City:", data.Locality)
	}
}
