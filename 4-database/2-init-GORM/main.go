package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primarykey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 100.00,
	// })

	products := []Product{
		{Name: "Notebook", Price: 100.00},
		{Name: "Mouse", Price: 50.00},
		{Name: "Keyboard", Price: 90.00},
	}

	db.Create(&products)
}
