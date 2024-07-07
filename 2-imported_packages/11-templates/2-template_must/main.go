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

func main() {

	course := Course{"Go", 49}
	tmp := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} - Workload: {{.Workload}}"))

	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
