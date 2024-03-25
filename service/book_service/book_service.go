package book_service

import (
	"reserve/models"
	"reserve/models/book"
)

func NewBookService() IBookService {
	return &BookService{}
}

type IBookService interface {
	Create(in *book.AddBookRequest) error
	List(in *book.ListBookRequest) ([]*book.ListBookRespond, error)
}

type BookService struct{}

func (b *BookService) List(in *book.ListBookRequest) ([]*book.ListBookRespond, error) {
	resp := make([]*book.ListBookRespond, 0)
	date, err := book.NewIBook().List(in)
	if err != nil {
		return resp, err
	}
	for _, v := range date {
		resp = append(resp, &book.ListBookRespond{
			Name:     v.Name,
			StoreNum: v.StoreNum,
			Price:    v.Price,
		})
	}
	return resp, nil
}

func (b *BookService) Create(in *book.AddBookRequest) error {
	insertDate := &models.Book{
		Name:     in.Name,
		StoreNum: in.StoreNum,
		Price:    in.Price,
	}
	return book.NewIBook().Create(insertDate)
}
