package main

import "fmt"

func main() {

	var myVar interface{} = "Gabreil rocha"
	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("The value  is %v an result is %v \n", res, ok)
}
