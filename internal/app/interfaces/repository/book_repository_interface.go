package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type BookRepositoryInterface interface {
	Get() (books []model.Book, apiError *errors.ApiError)
	Find(id string) (book model.Book, apiError *errors.ApiError)
	Create(book model.Book) (uint, *errors.ApiError)
	Update(id string, upBook model.Book) (model.Book, *errors.ApiError)
	Delete(id string) (apiError *errors.ApiError)
	BeforeCreate(categoriesId string) ([]uint, *errors.ApiError)
	AfterCreate(bookId uint, categoriesId []uint)
	BeforeUpdate(bookId uint, categoriesId string) *errors.ApiError
	BeforeDelete(bookId uint)
}
