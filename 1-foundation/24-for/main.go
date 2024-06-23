package main

func main() {
	numbers := []string{"one", "two", "three"}
	for k, v := range numbers {
		println(k, v)
	}

	for {
		println("Hello World")
	}
}
