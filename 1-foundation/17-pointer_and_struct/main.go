package main

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	myVar1 := 10
	myVar2 := 20
	sum(&myVar1, &myVar2)
	println(myVar1)
}
