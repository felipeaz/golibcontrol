package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews/model"
)

type ReviewRepositoryInterface interface {
	Get(bookId string) ([]model.Review, *errors.ApiError)
	Find(id string) (model.Review, *errors.ApiError)
	Create(review model.Review) (*model.Review, *errors.ApiError)
	Update(id string, upReview model.Review) *errors.ApiError
	Delete(id string) *errors.ApiError
}
