package main

import "net/http"

type Blog struct {
	title string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	// mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
	// 		w.Write([]byte("Hello World!"))
	// } )
	mux.Handle("/blog", Blog{title: "My Blog"})

	http.ListenAndServe("8080", mux)

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
