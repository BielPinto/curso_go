package main

import (
	"github.com/BielPinto/curso_go/5-packaging/4-go-workspaces/math"
)

// this comando work wiht packege local,good practices between packages
// raiz\# go work init ./math ./system
// go run system/main.go

// system/# go mod tidy -e  - >>  download the packages and ignore the packages that gave error
func main() {

	m := math.NewMath(1, 2)
	println(m.ADD())
	println(uuid.nNew().String())

}
