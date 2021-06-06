package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

type BookCategoryRepositoryInterface interface {
	GetCategoriesByIds(categoriesIds []uint) (categories []uint, apiError *errors.ApiError)
	CreateCategories(bookId uint, categoriesIds []uint)
	DeleteCategories(bookId uint)
}
