package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Notebook", 1899.90)

	err = inserProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 100.0

	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProduct(db, &product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %v, possui o preço de %.2f id: %v \n", p.Name, p.Price, p.ID)
}

func inserProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	// stmt, err := db.Prepare("insert into products(id, name, price) values($1, $2, $3)") db sqlit
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil

}

func updateProduct(db *sql.DB, product *Product) error {

	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectProduct(db *sql.DB, id *string) (*Product, error) {

	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return nil, err
	}

	return &p, nil

}
