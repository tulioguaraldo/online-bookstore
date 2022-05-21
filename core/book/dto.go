package book

import (
	"time"

	"github.com/tulioguaraldo/online-bookstore/model"
)

type BookResponse struct {
	Id          uint64    `json:"Id"`
	Name        string    `json:"BookName"`
	Author      string    `json:"BookAuthor"`
	Publisher   string    `json:"BookPublisher"`
	Price       float64   `json:"BookPrice"`
	WroteAt     time.Time `json:"WroteAt"`
	PublishedAt time.Time `json:"PublishedAt"`
}

func BookToResponse(book *model.Book, res *BookResponse) {
	*res = BookResponse{
		Id:          book.Id,
		Name:        book.Name,
		Author:      book.Author,
		Publisher:   book.Publisher,
		Price:       book.Price,
		WroteAt:     book.WroteAt,
		PublishedAt: book.WroteAt,
	}
}
