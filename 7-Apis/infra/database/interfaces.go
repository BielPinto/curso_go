package database

import "github.com/BielPinto/curso_go/7-Apis/internal/entity"

type UserInteface interface {
	Creat(User *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
