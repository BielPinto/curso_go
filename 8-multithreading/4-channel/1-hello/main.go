package main

import "fmt"

// Thread 1
func main() {
	channal1 := make(chan string) // is empty
	// Thread 2
	go func() {
		channal1 <- "Helo world" //is full
	}()

	// Thread 3
	msg := <-channal1 //is emptying
	fmt.Println(msg)
}
