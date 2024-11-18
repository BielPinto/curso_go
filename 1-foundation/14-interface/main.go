package main

import "fmt"

type address struct {
	Public_place string
	Number       int
	City         string
	State        string
}

type Person interface { //Interface can only pass methods and not properties
	Disable() //Ther signature must be identified in the methods used in the struct id Disable
}

type Company struct {
	name string
}

func (e Company) Disable() {}

type Client struct {
	Name   string
	Age    int
	Active bool
	address
}

func (c Client) Disable() {
	c.Active = false
	fmt.Printf("O client %s was disable %d \n", c.Name, c.Active)
}

func Dectivation(person Person) {
	person.Disable()
}

func main() {
	gabriel := Client{
		Name:   "Gabriel",
		Age:    30,
		Active: true,
	}

	Dectivation(gabriel)
	myCompany := Company{}
	Dectivation(myCompany)

}
