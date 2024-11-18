package main

func main() {
	//memory => address => value
	//&address * type pointer
	a := 10
	var pointer *int = &a
	*pointer = 20
	b := &a
	println(a)
	*b = 30
	println(a)
	println(&b)
	println(*b)
	println(*pointer)
	println(pointer)
	println(&pointer)
}
