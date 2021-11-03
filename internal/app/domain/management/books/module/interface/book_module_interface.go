package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model"
)

type BookModuleInterface interface {
	Get() ([]model.Book, *errors.ApiError)
	Find(id string) (model.Book, *errors.ApiError)
	Create(book model.Book) (uint, *errors.ApiError)
	Update(id string, upBook model.Book) *errors.ApiError
	Delete(id string) *errors.ApiError
}
