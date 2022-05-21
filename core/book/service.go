package book

import "github.com/tulioguaraldo/online-bookstore/model"

type InterfaceBookService interface {
	getAll() ([]model.Book, error)
	show(book *model.Book) (*model.Book, error)
	create(req *model.Book) (*model.Book, error)
}

type bookService struct {
	repository InterfaceBookRepository
}

func NewBookService(repository InterfaceBookRepository) bookService {
	return bookService{
		repository: repository,
	}
}

func (s *bookService) getAll() ([]model.Book, error) {
	return s.repository.getAll()
}

func (s *bookService) show(book *model.Book) (*model.Book, error) {
	return s.repository.show(book)
}

func (s *bookService) create(req *model.Book) (*model.Book, error) {
	return s.repository.create(req)
}
