package database

import "github.com/BielPinto/curso_go/7-Apis/internal/entity"

type UserInteface interface {
	Create(User *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductIterface interface {
	Create(product *entity.Product) error
	FindAll(paga, limit int, sort string) ([]entity.Product, error)
	FindByID(id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error
}
