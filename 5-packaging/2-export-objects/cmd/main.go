package main

import (
	"fmt"

	"github.com/BielPinto/curso_go/5-packaging/2-export-objects/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m)
	fmt.Println(m.C)
	fmt.Println(m.ADD())
	fmt.Println(math.X)
}
