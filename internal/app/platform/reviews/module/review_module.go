package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reviews"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/logger"
	"strconv"
)

type ReviewModule struct {
	Repository reviews.Repository
	Log        logger.LogInterface
}

func NewReviewModule(repo reviews.Repository, log logger.LogInterface) ReviewModule {
	return ReviewModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ReviewModule) Get(bookId string) ([]reviews.Review, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m ReviewModule) Find(id string) (reviews.Review, *errors.ApiError) {
	review, apiError := m.Repository.Find(id)
	if apiError != nil {
		return reviews.Review{}, apiError
	}
	bookReviews, apiError := m.Get(strconv.Itoa(int(review.BookId)))
	if apiError != nil {
		return reviews.Review{}, apiError
	}
	review.CalculateAvg(bookReviews)
	return review, nil
}

func (m ReviewModule) Create(comment reviews.Review) (*reviews.Review, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ReviewModule) Update(id string, upReview reviews.Review) *errors.ApiError {
	return m.Repository.Update(id, upReview)
}

func (m ReviewModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
