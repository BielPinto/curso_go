package main

import (
	"fmt"
	"mathematics/mathematics"
)

func main() {

	s := mathematics.Sum(10, 20)
	car := mathematics.Car{Brand: "fiat"}
	fmt.Println(car.Walk())
	fmt.Println("Result: %v", s)
}

//go mod init github.com/BielPinto/curso_go
