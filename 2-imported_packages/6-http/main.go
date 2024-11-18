package main

import "net/http"

func main() {

	http.HandleFunc("/", SearchZipCode)
	http.ListenAndServe(":8080", nil)

}

func SearchZipCode(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello World"))

}
