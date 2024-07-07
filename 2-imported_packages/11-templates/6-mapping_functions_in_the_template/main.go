package main

import (
	"net/http"
	"strings"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
	// Active   bool
}

type Courses []Course

func main() {

	http.HandleFunc("/", humeHanlder)
	http.ListenAndServe(":8282", nil)
}
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func humeHanlder(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(w, Courses{
		{"Go", 40},
		{"node", 250},
		{"Next", 75},
		{"Docker", 87},
	})
	if err != nil {
		panic(err)
	}
}
