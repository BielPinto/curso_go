package main

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	myvar1 := 10
	myvar2 := 10
	sum(&myvar1, &myvar2)
	println(myvar1)
}
