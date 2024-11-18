package main

import (
	"log"
	"net/http"
)

func main() {
	fileServe := http.FileServer(http.Dir("./public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServe)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("My Blog"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))

}
