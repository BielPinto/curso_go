package main

import "fmt"

func main() {
	fmt.Println(sum(1, 49, 5, 434, 434, 4))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
