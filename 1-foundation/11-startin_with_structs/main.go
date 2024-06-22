package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	gabriel := Client{
		Nome:  "gabriel",
		Idade: 30,
		Ativo: true,
	}
	fmt.Printf("Nome: %s, Idade%d, Ativo: %t \n", gabriel.Nome, gabriel.Idade, gabriel.Ativo)
	fmt.Println(gabriel.Ativo)
}
