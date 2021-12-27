package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model"
)

type BookRepositoryInterface interface {
	Get() (books []model.Book, apiError *errors.ApiError)
	GetByFilter(filter model.Filter) (books []model.Book, apiError *errors.ApiError)
	Find(id string) (book model.Book, apiError *errors.ApiError)
	Create(book model.Book) (*model.Book, *errors.ApiError)
	Update(id string, upBook model.Book) *errors.ApiError
	Delete(id string) (apiError *errors.ApiError)
}
