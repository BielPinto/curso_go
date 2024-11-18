package main

import "fmt"

func receive(name string, hello chan<- string) {
	hello <- name
}
func read(data <-chan string) {
	fmt.Println(<-data)
}

func main() {
	hello := make(chan string)
	go receive("Hello", hello)
	read(hello)
}
