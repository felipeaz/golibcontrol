package module

import (
	"net/http"
	"testing"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/category/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/category/repository/mock"
	"github.com/stretchr/testify/assert"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func TestGetCategory(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	m := CategoryModule{Repository: repository}

	// Exec
	categories, apiError := m.Get()

	// Validation
	assert.Nil(t, apiError)
	assert.NotNil(t, categories)
	assert.Equal(t, 5, int(categories[0].ID))
	assert.Equal(t, "Sci-Fi", categories[0].Name)
}

func TestGetCategoryError(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestError = true
	m := CategoryModule{Repository: repository}

	// Exec
	categories, apiError := m.Get()

	// Validation
	assert.Nil(t, categories)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestFindCategory(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	m := CategoryModule{Repository: repository}
	id := "5"

	// Exec
	category, apiError := m.Find(id)

	// Validation
	assert.Nil(t, apiError)
	assert.NotNil(t, category)
	assert.Equal(t, 5, int(category.ID))
	assert.Equal(t, "Sci-Fi", category.Name)
}

func TestFindCategoryError(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestError = true
	m := CategoryModule{Repository: repository}
	id := "5"

	// Exec
	category, apiError := m.Find(id)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, model.Category{}, category)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestFindCategoryNotFoundError(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestNotFoundError = true
	m := CategoryModule{Repository: repository}
	id := "5"

	// Exec
	category, apiError := m.Find(id)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, model.Category{}, category)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "category not found", apiError.Error)
}

func TestCreateCategory(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	m := CategoryModule{Repository: repository}
	category := model.Category{
		Name: "Sci-Fi",
	}

	// Exec
	id, apiError := m.Create(category)

	// Validation
	assert.Nil(t, apiError)
	assert.NotNil(t, id)
	assert.Equal(t, 5, int(id))
}

func TestCreateCategoryWithError(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestError = true
	m := CategoryModule{Repository: repository}
	category := model.Category{
		Name: "Sci-Fi",
	}

	// Exec
	id, apiError := m.Create(category)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(id))
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestCategoryUpdate(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	m := CategoryModule{Repository: repository}
	id := "5"
	upCategory := model.Category{
		Name: "Sci-Fi Updated",
	}

	// Exec
	category, apiError := m.Update(id, upCategory)

	// Validation
	assert.Nil(t, apiError)
	assert.NotEqual(t, model.Category{}, category)
	assert.Equal(t, upCategory.Name, category.Name)
}

func TestUpdateCategoryNotFound(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestNotFoundError = true
	m := CategoryModule{Repository: repository}
	id := "5"
	upCategory := model.Category{
		Name: "Sci-Fi Updated",
	}

	// Exec
	category, apiError := m.Update(id, upCategory)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, model.Category{}, category)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "category not found", apiError.Error)
}

func TestUpdateCategoryWithError(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestError = true
	m := CategoryModule{Repository: repository}
	id := "5"
	upCategory := model.Category{
		Name: "Sci-Fi Updated",
	}

	// Exec
	category, apiError := m.Update(id, upCategory)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, model.Category{}, category)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestUpdateCategoryError(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestUpdateError = true
	m := CategoryModule{Repository: repository}
	id := "5"
	upCategory := model.Category{
		Name: "Sci-Fi Updated",
	}

	// Exec
	category, apiError := m.Update(id, upCategory)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, model.Category{}, category)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestDeleteCategory(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	m := CategoryModule{Repository: repository}
	id := "5"

	// Exec
	apiError := m.Delete(id)

	// Validation
	assert.Nil(t, apiError)
}

func TestDeleteCategoryNotFound(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestNotFoundError = true
	m := CategoryModule{Repository: repository}
	id := "5"

	// Exec
	apiError := m.Delete(id)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.DeleteFailMessage, apiError.Message)
	assert.Equal(t, "category not found", apiError.Error)
}

func TestDeleteCategoryWithError(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestError = true
	m := CategoryModule{Repository: repository}
	id := "5"

	// Exec
	apiError := m.Delete(id)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.DeleteFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestDeleteCategoryError(t *testing.T) {
	// Init
	repository := mock.CategoryRepositoryMock{}
	repository.TestDeleteError = true
	m := CategoryModule{Repository: repository}
	id := "5"

	// Exec
	apiError := m.Delete(id)

	// Validation
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.DeleteFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}
