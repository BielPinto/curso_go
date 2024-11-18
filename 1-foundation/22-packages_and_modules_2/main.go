package main

import (
	"fmt"

	"curso_go/mathematic"
)

func main() {

	s := mathematic.Sum(10, 20)
	fmt.Println("Result: %v", s)
}

//go mod init github.com/BielPinto/curso_go
