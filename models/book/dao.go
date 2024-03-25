package book

import (
	"gorm.io/gorm"
	"reserve/models"
)

func NewIBook() IBook {
	return &Book{
		db: models.GetDb(),
	}
}

type IBook interface {
	Create(in *models.Book) error
	List(in *ListBookRequest) ([]*models.Book, error)
}

type Book struct {
	db *gorm.DB
}

func (b *Book) List(in *ListBookRequest) ([]*models.Book, error) {
	date := make([]*models.Book, 0)
	err := b.db.Find(&date).Error
	return date, err
}

func (b *Book) Create(in *models.Book) error {
	return b.db.Create(in).Error
}
