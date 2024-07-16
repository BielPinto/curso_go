package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
-Client.go must make an HTTP request to server.go requesting the dollar quote.

- pply the context on the client and a 300 ms timeout of the server.go response

- save the current quote in a file "quotacao.txt" in Dollar format:{value}

- 3 context return the log error if the execution time is insufficient
*/
type Quotation struct {
	// Code       string `json:"code"`
	// Codein     string `json:"codein"`
	// Name       string `json:"name"`
	// High       string `json:"high"`
	// Low        string `json:"low"`
	// VarBid     string `json:"varBid"`
	// PctChange  string `json:"pctChange"`
	Bid string `json:"bid"`
	// Ask        string `json:"ask"`
	// Timestamp  string `json:"timestamp"`
	// CreateDate string `json:"create_date"`

}

type TemplateData struct {
	DollarValues []string
}

func main() {
	fielisexists()
	http.HandleFunc("/", humeHanlder)
	http.HandleFunc("/cotacao", quotationHanlder)
	fmt.Printf("Client started  Port:%d \n", 8282)
	log.Fatal(http.ListenAndServe(":8282", nil))

}

func fielisexists() {
	f, erre := os.Create("cotacao.txt")
	if erre != nil {
		panic(erre)
	}
	f.Close()
}

func quotationHanlder(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()
	quotation, err := getQuotation(ctx)
	if err != nil {
		fmt.Printf("error when getting quote error: %s \n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = saveQuotationToFile(quotation)
	if err != nil {
		fmt.Printf("Unable to insert quote data into the file error: %s \n", err.Error())
		http.Error(w, "Unable to insert quote data into the file error", http.StatusInternalServerError)
		return
	}
	jsonBytes, err := json.Marshal(quotation)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	fmt.Printf("Dollar received successfully: %s \n", string(jsonBytes))
	w.Write(jsonBytes)
	// humeHanlder(w, r)
}

func humeHanlder(w http.ResponseWriter, r *http.Request) {
	cotacaoFile, err := readDollarValues("cotacao.txt")
	if err != nil {
		log.Fatalf("Failed to read dollar value: %v", err)
	}

	if len(cotacaoFile) == 0 {
		fmt.Printf("No dollar in file - > go to route http://localhost:8282/cotacao to get the quote \n")

	}
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	data := TemplateData{
		DollarValues: cotacaoFile,
	}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func readDollarValues(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("file error: %s \n", err.Error())
		return nil, err
	}
	defer file.Close()

	var dollarValues []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Dólar:") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				dollarValues = append(dollarValues, strings.TrimSpace(parts[1]))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %s \n", err.Error())
		return nil, err
	}

	return dollarValues, nil
}

func saveQuotationToFile(quotation Quotation) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data := fmt.Sprintf("Dólar: %s\n",
		quotation.Bid)

	_, err = file.WriteString(data)
	return err
}

func getQuotation(ctx context.Context) (Quotation, error) {

	// c := http.Client{Timeout: 200 * time.Millisecond}
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return Quotation{}, err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return Quotation{}, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return Quotation{}, err
	}

	if resp.StatusCode != http.StatusOK {

		return Quotation{}, fmt.Errorf(string(body))
	}

	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		return Quotation{}, err
	}

	return quotation, nil
}
