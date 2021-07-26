package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/model"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/review/repository/interface"
)

type ReviewModule struct {
	Repository _interface.ReviewRepositoryInterface
}

func NewReviewModule(repo _interface.ReviewRepositoryInterface) ReviewModule {
	return ReviewModule{
		Repository: repo,
	}
}

func (m ReviewModule) Get(bookId string) ([]model.Review, *errors.ApiError) {
	return m.Repository.Get(bookId)
}

func (m ReviewModule) Find(id string) (model.Review, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ReviewModule) Create(comment model.Review) (uint, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m ReviewModule) Update(id string, upReview model.Review) *errors.ApiError {
	return m.Repository.Update(id, upReview)
}

func (m ReviewModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
