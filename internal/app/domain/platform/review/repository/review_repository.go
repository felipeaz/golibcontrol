package repository

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/model/converter"
)

const (
	ServiceName = "PlatformService"
)

type ReviewRepository struct {
	DB database.GORMServiceInterface
}

func NewReviewRepository(db database.GORMServiceInterface) ReviewRepository {
	return ReviewRepository{
		DB: db,
	}
}

func (r ReviewRepository) Get(bookId string) ([]model.Review, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&[]model.Review{}, "book_id", bookId)
	if err != nil {
		return nil, &errors.ApiError{
			Service: ServiceName,
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	reviews, apiError := converter.ConvertToSliceReviewObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return reviews, nil
}

func (r ReviewRepository) Find(id string) (model.Review, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Review{}, id)
	if err != nil {
		return model.Review{}, &errors.ApiError{
			Service: ServiceName,
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	review, apiError := converter.ConvertToReviewObj(result)
	if apiError != nil {
		return model.Review{}, apiError
	}
	return review, nil
}

func (r ReviewRepository) Create(review model.Review) (uint, *errors.ApiError) {
	err := r.DB.Persist(&review)
	if err != nil {
		return 0, &errors.ApiError{
			Service: ServiceName,
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return review.ID, nil
}

func (r ReviewRepository) Update(id string, upReview model.Review) *errors.ApiError {
	err := r.DB.Refresh(&upReview, id)
	if err != nil {
		return &errors.ApiError{
			Service: ServiceName,
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r ReviewRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&model.Review{}, id)
	if err != nil {
		return &errors.ApiError{
			Service: ServiceName,
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
