package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

/*
- Server.go must consume the API containing the Dollar and Real exchange rate at the address: nd then it should return in JSON format  - > done

- endpoint /quotation and port 8080 -> done

- use context with call to
api quote with 200ms timeout
and the inset in the database with a timeout of 10 ms.  -> done

- return only the bind field in the api  -> done
- record each quote received in the SQLite database driver _ "github.com/mattn/go-sqlite3"
- 3 context return the log error if the execution time is insufficient  -> doing
*/

type Quotation struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	// mux := http.NewServeMux()

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	http.HandleFunc("/cotacao", HandlerQuotation(db))
	http.ListenAndServe(":8080", nil)
}

func HandlerQuotation(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancel()
		quotation, err := getQuotation(ctx, "USD-BRL")
		if err != nil {
			fmt.Println("err", err)
			http.Error(w, "Unable to obtain Quotation", http.StatusInternalServerError)
			return
		}
		err = inserProduct(db, quotation)
		// err = saveQuotationToFile(quotation)
		if err != nil {
			fmt.Println("err", err)
			http.Error(w, "Unable to sbe quote to file", http.StatusInternalServerError)
			return
		}

		bid := struct {
			Bid string `json:"bid"`
		}{
			Bid: quotation.USDBRL.Bid,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(bid)
		// err = json.NewEncoder(w).Encode(quotation)

		if err != nil {
			http.Error(w, "Erro ao codificar a resposta JSON", http.StatusInternalServerError)
			return
		}
	}
}

func getQuotation(ctx context.Context, cod string) (Quotation, error) {

	// c := http.Client{Timeout: 200 * time.Millisecond}
	url := fmt.Sprintf("https://economia.awesomeapi.com.br/json/last/%s", cod)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return Quotation{}, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Quotation{}, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("outil.ReadAll( error", err)
		return Quotation{}, err
	}
	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		fmt.Println("Unmarshal(", err)
		return Quotation{}, err
	}
	return quotation, nil
}

func saveQuotationToFile(quotation Quotation) error {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	data := fmt.Sprintf("Code: %s, Codein: %s, Name: %s, High: %s, Low: %s, VarBid: %s, PctChange: %s, Bid: %s, Ask: %s, Timestamp: %s, CreateDate: %s\n",
		quotation.USDBRL.Code, quotation.USDBRL.Codein, quotation.USDBRL.Name, quotation.USDBRL.High, quotation.USDBRL.Low, quotation.USDBRL.VarBid, quotation.USDBRL.PctChange, quotation.USDBRL.Bid, quotation.USDBRL.Ask, quotation.USDBRL.Timestamp, quotation.USDBRL.CreateDate)

	_, err = file.WriteString(data)
	return err
}

func inserProduct(db *sql.DB, quotation Quotation) error {
	stmt, err := db.Prepare("insert into products(id, Code, Codein, Name, High, Low, VarBid, PctChange, Bid, Ask, CreateDate) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid.NewString(), quotation.USDBRL.Code, quotation.USDBRL.Codein, quotation.USDBRL.Name, quotation.USDBRL.High, quotation.USDBRL.Low, quotation.USDBRL.VarBid, quotation.USDBRL.PctChange, quotation.USDBRL.Bid, quotation.USDBRL.Ask, quotation.USDBRL.CreateDate)
	if err != nil {
		return err
	}
	return nil

}
