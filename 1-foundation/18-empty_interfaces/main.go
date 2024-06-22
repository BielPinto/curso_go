package main

import "fmt"

func main() {

	var x interface{} = 10
	var y interface{} = "Hello, world!"
	showType(x)
	showType(y)

}

func showType(t interface{}) {
	fmt.Println("The Type of the variabel is %T and the values is %v \n", t, t)
}
