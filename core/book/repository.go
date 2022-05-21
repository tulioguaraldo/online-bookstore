package book

import (
	"github.com/tulioguaraldo/online-bookstore/model"
	"gorm.io/gorm"
)

type InterfaceBookRepository interface {
	getAll() ([]model.Book, error)
	show(book *model.Book) (*model.Book, error)
	create(req *model.Book) (*model.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) bookRepository {
	return bookRepository{
		db: db,
	}
}

func (r *bookRepository) getAll() ([]model.Book, error) {
	var books []model.Book

	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (r *bookRepository) show(book *model.Book) (*model.Book, error) {
	if err := r.db.First(&book, &book.Id).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (r *bookRepository) create(req *model.Book) (*model.Book, error) {
	if err := r.db.Create(&req).Error; err != nil {
		return nil, err
	}

	return req, nil
}
