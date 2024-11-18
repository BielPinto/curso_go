package main

import "fmt"

func main() {

	s := []int{10, 20, 30, 50, 60, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])
	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v", len(s), cap(s), s)
	s = append(s, 120)
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
	s = append(s, 150, 44, 44, 44, 44, 44, 55, 44)
	fmt.Printf("len=%d cap=%d %v \n", len(s), cap(s), s)
}
