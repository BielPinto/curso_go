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

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("O client %s was disable %d \n", c.Name, c.Active)
}

func main() {
	gabriel := Client{
		Name:   "Gabriel",
		Age:    30,
		Active: true,
	}

	fmt.Printf("Name: %s, Age: %d, Active: %t \n", gabriel.Name, gabriel.Age, gabriel.Active)
	gabriel.Disable()
	fmt.Printf("Name: %s, Age: %d, Active: %t \n", gabriel.Name, gabriel.Age, gabriel.Active)

}
