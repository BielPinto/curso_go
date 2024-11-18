package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request Started")
	defer log.Println("Request finished")
	select {
	case <-time.After(time.Second * 5):
		log.Println("Request processed successfully")
		w.Write([]byte("Request processed successfully\n!"))
	case <-ctx.Done():
		log.Println("Request canceled by client")
	}

}
