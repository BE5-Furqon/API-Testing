package repository

import (
	"layerarch/config"
	"layerarch/entities"
	"layerarch/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookRepo(t *testing.T)  {
	config := config.GetConfig()
	db := util.InitDB(config)

	db.Migrator().DropTable(&entities.Book{})
	db.AutoMigrate(&entities.Book{})

	repo := NewBookRepo(db)
	
	mockUser := entities.Book{
		Title: "Clean Code",
		Author: "Martin",
	}

	t.Run("Insert", func(t *testing.T) {
		res, err := repo.Insert(mockUser)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.Title, res.Title)
		assert.Equal(t, mockUser.Author, res.Author)
	})
	
	t.Run("Get All", func(t *testing.T) {
		res, err := repo.GetAll()

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.Title, res[0].Title)
		assert.Equal(t, mockUser.Author, res[0].Author)
	})

	t.Run("Get", func(t *testing.T) {
		res, err := repo.Get(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.Title, res.Title)
		assert.Equal(t, mockUser.Author, res.Author)
	})
	
	t.Run("Edit", func(t *testing.T) {
		res, err := repo.Edit(mockUser, 1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, mockUser.Title, res.Title)
		assert.Equal(t, mockUser.Author, res.Author)
	})
	
	t.Run("Delete", func(t *testing.T) {
		res, err := repo.Delete(1)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})
	
}