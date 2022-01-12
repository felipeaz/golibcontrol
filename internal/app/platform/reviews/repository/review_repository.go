package repository

import (
	"fmt"
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
	tx := r.DB.Where(nil, fmt.Sprintf("book_id = %s", bookId))
	result, err := r.DB.Find(tx, &[]reviews.Review{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &reviews.Review{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upReview)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &reviews.Review{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
