package main

import (
	"fmt"

	"github.com/BielPinto/curso_go/5-packaging/1-introduction/math"
	// "meuModulo/math"
)

// whenever you can import the modules via the url
func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.ADD())
	// fmt.Println("Hello, world!")
}
