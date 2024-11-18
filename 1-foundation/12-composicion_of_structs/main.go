package main

import "fmt"

type Addres struct {
	Public_place string
	Number       int
	City         string
	State        string
}

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	gabriel := Client{
		Name:   "Gabriel",
		Age:    30,
		Active: true,
	}
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", gabriel.Name, gabriel.Age, gabriel.Active)
}
