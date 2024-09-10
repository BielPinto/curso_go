package main

import (
	"fmt"
	"net/http"
)

// golang 1.22
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /books/{id}", GeTBookHandler)
	mux.HandleFunc("GET /books/dir/{d...}", BooksPathHandler)
	mux.HandleFunc("GET /books/{$}", BooksHandler) //exato
	mux.HandleFunc("GET /books/precedence/latest", BooksPrecedenceHanlder)
	mux.HandleFunc("GET /books/precedence/{x}", BooksPrecedence2Hanlder)
	mux.HandleFunc("GET /books/{id}", BooksPrecedenceHanlder)
	mux.HandleFunc("GET /{x}/latest", BooksPrecedence2Hanlder)

	http.ListenAndServe(":9000", mux)

}

func GeTBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Book" + id))
}

func BooksPathHandler(w http.ResponseWriter, r *http.Request) {
	dirparth := r.PathValue("d") // Access captured directory path segments as
	fmt.Fprintf(w, "Accessing directory path: %s \n", dirparth)
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books"))
}
func BooksPrecedenceHanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence"))
}

func BooksPrecedence2Hanlder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Books Precedence 2 "))
}
