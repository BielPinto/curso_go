package main

import "fmt"

func main() {
	total := func() int {
		return sum(1, 49, 5, 434, 434, 4) * 2
	}()

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
