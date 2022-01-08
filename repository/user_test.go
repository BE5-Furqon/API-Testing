package repository

import (
	"layerarch/config"
	"layerarch/entities"
	"layerarch/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepo(t *testing.T)  {
	config := config.GetConfig()
	db := util.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := NewUserRepo(db)
	
	mockUser := entities.User{
		Name:           "Furqon",
		Email:          "furqonzt99@gmail.com",
		Password:       "1234qwer",
	}

	t.Run("Register", func(t *testing.T) {
		res, err := repo.Register(mockUser)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.Name, res.Name)
		assert.Equal(t, mockUser.Email, res.Email)
		assert.Equal(t, mockUser.Password, res.Password)
	})
	
	t.Run("Login", func(t *testing.T) {
		res, err := repo.Login(mockUser.Email, mockUser.Password)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
	
	t.Run("Get All", func(t *testing.T) {
		res, err := repo.GetAll()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.Name, res[0].Name)
		assert.Equal(t, mockUser.Email, res[0].Email)
		assert.Equal(t, mockUser.Password, res[0].Password)
	})

	t.Run("Get", func(t *testing.T) {
		res, err := repo.Get(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.Name, res.Name)
		assert.Equal(t, mockUser.Email, res.Email)
		assert.Equal(t, mockUser.Password, res.Password)
	})
	
	t.Run("Edit", func(t *testing.T) {
		res, err := repo.Edit(mockUser, 1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.Name, res.Name)
		assert.Equal(t, mockUser.Email, res.Email)
		assert.Equal(t, mockUser.Password, res.Password)
	})
	
	t.Run("Delete", func(t *testing.T) {
		res, err := repo.Delete(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
	
}