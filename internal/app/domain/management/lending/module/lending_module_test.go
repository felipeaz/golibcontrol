package module

import (
	"net/http"
	"testing"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/repository/mock"
	"github.com/stretchr/testify/assert"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func TestGetLending(t *testing.T) {
	r := mock.LendingRepositoryMock{}

	m := LendingModule{Repository: r}

	lendings, apiError := m.Get()

	assert.Nil(t, apiError)
	assert.NotNil(t, lendings)
	assert.Equal(t, 25, int(lendings[0].ID))
	assert.Equal(t, 5, int(lendings[0].BookID))
	assert.Equal(t, 10, int(lendings[0].StudentID))
	assert.Equal(t, time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC), lendings[0].LendingDate)
}

func TestGetLendingError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestError = true

	m := LendingModule{Repository: r}

	lendings, apiError := m.Get()

	assert.Nil(t, lendings)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestFindLending(t *testing.T) {
	r := mock.LendingRepositoryMock{}

	m := LendingModule{Repository: r}
	id := "25"

	lending, apiError := m.Find(id)

	assert.Nil(t, apiError)
	assert.NotNil(t, lending)
	assert.Equal(t, 25, int(lending.ID))
	assert.Equal(t, 5, int(lending.BookID))
	assert.Equal(t, 10, int(lending.StudentID))
	assert.Equal(t, time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC), lending.LendingDate)
}

func TestFindLendingError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestError = true

	m := LendingModule{Repository: r}
	id := "25"

	lending, apiError := m.Find(id)

	assert.NotNil(t, apiError)
	assert.Equal(t, model.Lending{}, lending)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestFindLendingNotFoundError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestNotFoundError = true

	m := LendingModule{Repository: r}
	id := "25"

	lending, apiError := m.Find(id)

	assert.NotNil(t, apiError)
	assert.Equal(t, model.Lending{}, lending)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "lending not found", apiError.Error)
}

func TestCreateLending(t *testing.T) {
	r := mock.LendingRepositoryMock{}

	m := LendingModule{Repository: r}
	lending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(lending)
	assert.Nil(t, apiError)
	assert.Equal(t, 25, int(id))
}

func TestCreateLendingStudentNotFoundError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestStudentNotFoundError = true

	m := LendingModule{Repository: r}
	lending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(lending)

	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(id))
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "student not found", apiError.Error)
}

func TestCreateLendingStudentError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestStudentNotFoundError = true
	r.TestError = true

	m := LendingModule{Repository: r}
	lending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(lending)

	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(id))
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestCreateLendingBookNotFoundError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestBookNotFoundError = true

	m := LendingModule{Repository: r}
	lending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(lending)

	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(id))
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "book not found", apiError.Error)
}

func TestCreateLendingBookError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestBookNotFoundError = true
	r.TestError = true

	m := LendingModule{Repository: r}
	lending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(lending)

	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(id))
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestCreateLendingBookAlreadyLentError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestBookAlreadyLentError = true

	m := LendingModule{Repository: r}
	lending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(lending)

	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(id))
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "book is already lent", apiError.Error)
}

func TestCreateLendingStudentAlreadyLentError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestStudentAlreadyLentError = true

	m := LendingModule{Repository: r}
	lending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(lending)

	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(id))
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "student has already lent a book", apiError.Error)
}

func TestCreateLendingError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestError = true

	m := LendingModule{Repository: r}
	lending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(lending)

	assert.NotNil(t, apiError)
	assert.Equal(t, 0, int(id))
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestUpdateLending(t *testing.T) {
	r := mock.LendingRepositoryMock{}

	m := LendingModule{Repository: r}
	id := "25"
	upLending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upLending)
	assert.Nil(t, apiError)
}

func TestUpdateLendingNotFoundError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestNotFoundError = true

	m := LendingModule{Repository: r}
	id := "25"
	upLending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upLending)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "lending not found", apiError.Error)

}

func TestUpdateLendingStudentNotFoundError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestStudentNotFoundError = true

	m := LendingModule{Repository: r}
	id := "25"
	upLending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upLending)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "student not found", apiError.Error)
}

func TestUpdateLendingStudentError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestStudentNotFoundError = true
	r.TestError = true

	m := LendingModule{Repository: r}
	id := "25"
	upLending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upLending)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestUpdateLendingBookNotFoundError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestBookNotFoundError = true

	m := LendingModule{Repository: r}
	id := "25"
	upLending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upLending)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "book not found", apiError.Error)
}

func TestUpdateLendingBookError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestBookNotFoundError = true
	r.TestError = true

	m := LendingModule{Repository: r}
	id := "25"
	upLending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upLending)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestUpdateLendingError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestUpdateError = true

	m := LendingModule{Repository: r}
	id := "25"
	upLending := model.Lending{
		ID:          25,
		BookID:      5,
		StudentID:   10,
		LendingDate: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upLending)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestDeleteLending(t *testing.T) {
	r := mock.LendingRepositoryMock{}

	m := LendingModule{Repository: r}
	id := "25"

	apiError := m.Delete(id)
	assert.Nil(t, apiError)
}

func TestDeleteLendingError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestDeleteError = true

	m := LendingModule{Repository: r}
	id := "25"

	apiError := m.Delete(id)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.DeleteFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestDeleteLendingNotFoundError(t *testing.T) {
	r := mock.LendingRepositoryMock{}
	r.TestNotFoundError = true

	m := LendingModule{Repository: r}
	id := "25"

	apiError := m.Delete(id)
	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.DeleteFailMessage, apiError.Message)
	assert.Equal(t, "lending not found", apiError.Error)
}
