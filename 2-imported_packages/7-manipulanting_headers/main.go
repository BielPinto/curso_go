package main

import "net/http"

func main() {
	http.HandleFunc("/", SearchZipCode)
	http.ListenAndServe(":8080", nil)
}

func SearchZipCode(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipCodeParam := r.URL.Query().Get("zip")
	if zipCodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("hello World"))
}
