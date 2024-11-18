package main

import "fmt"

func myPanic1() {
	panic("panic1")
}
func myPanic2() {
	panic("panic22")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == "panic1" {
				fmt.Println("panic1 ")
			}
			if r == "panic2" {
				fmt.Println("panic2 ")
			}

		}
	}()

	myPanic1()
	// myPanic2()
}
