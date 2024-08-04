package main

import "github.com/BielPinto/curso_go/5-packaging/3-go-mod/math"

// this comando work wiht packege local,
//go mod edit -replace github.com/BielPinto/curso_go/5-packaging/3-go-mod/math=../math
// go mod tidy

func main() {

	m := math.NewMath(1, 2)
	println(m.ADD())

}
