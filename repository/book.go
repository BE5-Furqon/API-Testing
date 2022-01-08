package repository

import (
	"layerarch/entities"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}

type Book interface {
	GetAll() ([]entities.Book, error)
	Get(bookId int) (entities.Book, error)
	Insert(entities.Book) (entities.Book, error)
	Edit(book entities.Book, bookId int) (entities.Book, error)
	Delete(bookId int) (entities.Book, error)
}

func (br *BookRepository) GetAll() ([]entities.Book, error) {
	var books []entities.Book

	if err := br.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (br *BookRepository) Get(bookId int) (entities.Book, error) {
	var book entities.Book

	if err := br.db.First(&book, bookId).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (br *BookRepository) Insert(book entities.Book) (entities.Book, error) {
	if err := br.db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (br *BookRepository) Edit(book entities.Book, bookId int) (entities.Book, error) {
	var b entities.Book
	br.db.First(&b, bookId)

	if err := br.db.Model(&b).Updates(book).Error; err != nil {
		return b, err
	}

	return b, nil
}

func (br *BookRepository) Delete(bookId int) (entities.Book, error) {
	var book entities.Book

	if err := br.db.Delete(&book, bookId).Error; err != nil {
		return book, err
	}

	return book, nil
}