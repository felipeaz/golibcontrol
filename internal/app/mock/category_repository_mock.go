package mock

import (
	"net/http"
	"strconv"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/model"
)

type CategoryRepositoryMock struct {
	TestError         bool
	TestNotFoundError bool
}

func (r CategoryRepositoryMock) Get() (categories []model.Category, apiError *errors.ApiError) {
	if r.TestError {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked error",
		}
	}

	categories = []model.Category{
		model.Category{
			ID:        5,
			Name:      "Sci-Fi",
			CreatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
			UpdatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		},
	}

	return categories, nil
}

func (r CategoryRepositoryMock) Find(id string) (category model.Category, apiError *errors.ApiError) {
	if r.TestError {
		return model.Category{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   "mocked error",
		}
	} else if r.TestNotFoundError {
		return model.Category{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   "category not found",
		}
	}

	category = model.Category{
		ID:        5,
		Name:      "Sci-Fi",
		CreatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	return category, nil
}

func (r CategoryRepositoryMock) Create(category model.Category) (uint, *errors.ApiError) {
	if r.TestError {
		return 0, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   "mocked error",
		}
	}

	return 5, nil
}

func (r CategoryRepositoryMock) Update(id string, upCategory model.Category) (model.Category, *errors.ApiError) {
	if r.TestNotFoundError {
		return model.Category{}, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.UpdateFailMessage,
			Error:   "category not found",
		}
	} else if r.TestError {
		return model.Category{}, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "mocked error",
		}
	}

	idInt, _ := strconv.Atoi(id)
	upCategory.ID = uint(idInt)

	return upCategory, nil
}

func (r CategoryRepositoryMock) Delete(id string) (apiError *errors.ApiError) {
	if r.TestNotFoundError {
		return &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.DeleteFailMessage,
			Error:   "category not found",
		}
	} else if r.TestError {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   "mocked error",
		}
	}

	return nil
}
