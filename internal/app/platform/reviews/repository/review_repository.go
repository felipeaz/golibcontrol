package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/reviews/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
)

type ReviewRepository struct {
	DB database.GORMServiceInterface
}

func NewReviewRepository(db database.GORMServiceInterface) ReviewRepository {
	return ReviewRepository{
		DB: db,
	}
}

func (r ReviewRepository) Get(bookId string) ([]reviews.Review, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&[]reviews.Review{}, "book_id", bookId)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceReviewObj(result)
}

func (r ReviewRepository) Find(id string) (reviews.Review, *errors.ApiError) {
	result, err := r.DB.Fetch(&reviews.Review{}, id)
	if err != nil {
		return reviews.Review{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToReviewObj(result)
}

func (r ReviewRepository) Create(review reviews.Review) (*reviews.Review, *errors.ApiError) {
	err := r.DB.Persist(&review)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &review, nil
}

func (r ReviewRepository) Update(id string, upReview reviews.Review) *errors.ApiError {
	err := r.DB.Refresh(&upReview, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r ReviewRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&reviews.Review{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
