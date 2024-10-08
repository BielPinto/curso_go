package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primarykey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})
	//create
	// db.Create(&Product{
	// 	Name:  "Mouse",
	// 	Price: 90.00,
	// })

	//create batch
	// products := []Product{
	// 	{Name: "Notebook", Price: 100.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 90.00},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// db.First(&product, 2)
	// fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	// var products []Product

	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//where
	// var products []Product
	// db.Where("price < ?", 90).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//like
	// var products []Product
	// db.Where("name like ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//update and delete
	// var p Product
	// db.First(&p, 9)
	// fmt.Println(p)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 9)
	// fmt.Println(p2)
	// db.Delete(&p2)

}
