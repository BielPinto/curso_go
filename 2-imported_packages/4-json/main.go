package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"`
	Balance int `json:"s"`
	// balance int `json:"-"`
	// balance int `json:"s" validate:"gt=0"`
}

func main() {

	account := Account{Number: 1, Balance: 100}
	//when I use marshal I save the json value for myself
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(res))
	/* When using  the encoderm you get the vakue and carry out the seelection process and deliver it to someone,
	   either to stdout or putting it in the file or webserve*/
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	jsonPure := []byte(`{"n":2,"s":200}`)
	var accountX Account
	json.Unmarshal(jsonPure, &accountX)

	println(accountX.Balance)
}
