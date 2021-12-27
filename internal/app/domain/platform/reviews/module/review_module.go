package module

import (
	"strconv"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews/model"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type ReviewModule struct {
	Repository _interface.ReviewRepositoryInterface
	Log        logger.LogInterface
}

func NewReviewModule(repo _interface.ReviewRepositoryInterface, log logger.LogInterface) ReviewModule {
	return ReviewModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ReviewModule) Get(bookId string) ([]model.Review, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m ReviewModule) Find(id string) (model.Review, *errors.ApiError) {
	review, apiError := m.Repository.Find(id)
	if apiError != nil {
		return model.Review{}, apiError
	}
	bookReviews, apiError := m.Get(strconv.Itoa(int(review.BookId)))
	if apiError != nil {
		return model.Review{}, apiError
	}
	review.CalculateAvg(bookReviews)
	return review, nil
}

func (m ReviewModule) Create(comment model.Review) (*model.Review, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ReviewModule) Update(id string, upReview model.Review) *errors.ApiError {
	return m.Repository.Update(id, upReview)
}

func (m ReviewModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
