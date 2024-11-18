package database

import (
	"testing"

	"github.com/BielPinto/curso_go/7-Apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreatUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memonry:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Gabriel", "j@j.com", "123456")
	userDb := NewUser(db)

	err = userDb.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("gabriel", "j@j.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	userFount, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFount.ID)
	assert.Equal(t, user.Name, userFount.Name)
	assert.Equal(t, user.Email, userFount.Email)
	assert.NotNil(t, userFount.Password)

}
