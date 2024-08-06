package main

import configs "github.com/BielPinto/curso_go/7-Apis/config"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)

}
