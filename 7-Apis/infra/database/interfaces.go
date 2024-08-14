package database

import "github.com/BielPinto/curso_go/7-Apis/internal/entity"

type UserInteface interface {
	Creat(User *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductIterface interface {
	Creat(product *entity.Product) error
	FindAll(paga, limit int, sort string) ([]entity.Product, error)
	UpdateById(id string) (*entity.Product, error)
	update(product *entity.Product) error
	Delete(id string) error
}
