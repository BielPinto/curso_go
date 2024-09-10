package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser

	jsonData := `{"user":{"name":"John Doe","age":30}}`

	value, err := p.Parse(jsonData)
	if err != nil {
		panic(err)
	}

	user := value.GetObject("user")
	fmt.Printf("User name: %s\n", user.Get("name"))
	fmt.Printf("User age: %s\n", user.Get("age"))

	userJSON := value.Get("user").String()
	var user1 User
	if err := json.Unmarshal([]byte(userJSON), &user1); err != nil {
		panic(err)
	}

	fmt.Println(user1.Name, user1.Age)
}
