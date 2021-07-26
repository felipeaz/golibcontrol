package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/model/converter"
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
	result, apiError := r.DB.FindWhere(&[]model.Review{}, "book_id", bookId)
	if apiError != nil {
		return nil, apiError
	}
	reviews, apiError := converter.ConvertToSliceReviewObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return reviews, nil
}

func (r ReviewRepository) Find(id string) (model.Review, *errors.ApiError) {
	result, apiError := r.DB.Find(&model.Review{}, id)
	if apiError != nil {
		return model.Review{}, apiError
	}
	review, apiError := converter.ConvertToReviewObj(result)
	if apiError != nil {
		return model.Review{}, apiError
	}
	return review, nil
}

func (r ReviewRepository) Create(review model.Review) (uint, *errors.ApiError) {
	apiError := r.DB.Create(&review)
	if apiError != nil {
		return 0, apiError
	}
	return review.ID, nil
}

func (r ReviewRepository) Update(id string, upReview model.Review) *errors.ApiError {
	return r.DB.Update(&upReview, id)
}

func (r ReviewRepository) Delete(id string) *errors.ApiError {
	return r.DB.Delete(&model.Review{}, id)
}
