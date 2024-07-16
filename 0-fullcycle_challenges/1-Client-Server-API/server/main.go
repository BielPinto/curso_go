package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

/*
- Server.go must consume the API containing the Dollar and Real exchange rate at the address: nd then it should return in JSON format  --> done
- endpoint /cotacao and port 8080 --> done
- record each quote received in the SQLite database driver _ "github.com/mattn/go-sqlite3" - > doing
- use context with call to api quote with 200ms timeout and the inset in the database with a timeout of 10 ms.  --> done
- return only the bind field in the api --> done
- 3 context return the log error if the execution time is insuficient  --> done
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
	db := bdExists()
	defer db.Close()
	http.HandleFunc("/cotacao", HandlerQuotation(db))

	fmt.Println("Server Up port:8080")
	http.ListenAndServe(":8080", nil)
}

func bdExists() (db *sql.DB) {
	dirPath := "./data"
	fileDB := "database.db"

	dbPath := fmt.Sprintf("%s/%s", dirPath, fileDB)

	err := createDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS quotation (
        id TEXT PRIMARY KEY,
        code TEXT,
        codein TEXT,
        name TEXT,
        high REAL,
        low REAL,
        varBid REAL,
        PctChange REAL,
        Bid REAL,
        Ask REAL,
        CreateDate TEXT
    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected and table created.")
	return db
}

func createDir(path string) error {

	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
		return err
	}
	return nil
}

func HandlerQuotation(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
		defer cancel()
		quotation, err := getQuotation(ctx, "USD-BRL")
		if err != nil {
			fmt.Printf("Unable to obtain Quotation err: %s \n", err)
			http.Error(w, "Unable to obtain Quotation", http.StatusInternalServerError)
			return
		}
		err = inserProduct(db, quotation)
		if err != nil {
			fmt.Printf("Unable to insert quote data into the database error: %s \n", err)
			http.Error(w, "Unable to insert quote data into the database", http.StatusInternalServerError)
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
			fmt.Printf("error when converting object to json error: %s \n", err)
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Quotation{}, err
	}
	var quotation Quotation
	err = json.Unmarshal(body, &quotation)
	if err != nil {
		return Quotation{}, err
	}
	return quotation, nil
}

func inserProduct(db *sql.DB, quotation Quotation) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	stmt, err := db.PrepareContext(ctx, "insert into quotation(id, Code, Codein, Name, High, Low, VarBid, PctChange, Bid, Ask, CreateDate) values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, uuid.NewString(), quotation.USDBRL.Code, quotation.USDBRL.Codein, quotation.USDBRL.Name, quotation.USDBRL.High, quotation.USDBRL.Low, quotation.USDBRL.VarBid, quotation.USDBRL.PctChange, quotation.USDBRL.Bid, quotation.USDBRL.Ask, quotation.USDBRL.CreateDate)
	id, _ := res.LastInsertId()
	fmt.Printf("Registered successfully! Id: %d \n", id)
	if err != nil {
		return err
	}
	return nil

}
