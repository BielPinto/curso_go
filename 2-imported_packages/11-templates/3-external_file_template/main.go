package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
	// Active   bool
}

type Courses []Course

func main() {

	t := template.Must(template.New("template.html").ParseFiles("template.html"))

	err := t.Execute(os.Stdout, Courses{
		{"Go", 40},
		{"node", 250},
		{"Next", 75},
		{"Docker", 87},
	})
	if err != nil {
		panic(err)
	}
}
