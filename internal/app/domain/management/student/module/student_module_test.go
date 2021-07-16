package module

import (
	"net/http"
	"testing"
	"time"

	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/repository/mock"
	"github.com/stretchr/testify/assert"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

func TestGetStudent(t *testing.T) {
	r := mock.StudentRepositoryMock{}

	m := StudentModule{Repository: r}

	students, apiError := m.Get()
	student := students[0]

	assert.Nil(t, apiError)
	assert.NotNil(t, students)
	assert.Equal(t, "2500651", student.ID)
	assert.Equal(t, "Felipe de Azevedo Silva", student.Name)
	assert.Equal(t, "felipe9_azevedo@hotmail.com", student.Email)
	assert.Equal(t, "(00)00000-0000", student.Phone)
	assert.Equal(t, "7th", student.Grade)
	assert.Equal(t, "31/12/1997", student.Birthday)
	assert.Equal(t, time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC), student.CreatedAt)
	assert.Equal(t, time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC), student.UpdatedAt)
}

func TestGetStudentError(t *testing.T) {
	r := mock.StudentRepositoryMock{}
	r.TestError = true

	m := StudentModule{Repository: r}

	students, apiError := m.Get()

	assert.NotNil(t, apiError)
	assert.Nil(t, students)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestFind(t *testing.T) {
	r := mock.StudentRepositoryMock{}

	m := StudentModule{Repository: r}
	id := "25"

	student, apiError := m.Find(id)

	assert.Nil(t, apiError)
	assert.NotEqual(t, model.Student{}, student)
	assert.Equal(t, "2500651", student.ID)
	assert.Equal(t, "Felipe de Azevedo Silva", student.Name)
	assert.Equal(t, "felipe9_azevedo@hotmail.com", student.Email)
	assert.Equal(t, "(00)00000-0000", student.Phone)
	assert.Equal(t, "7th", student.Grade)
	assert.Equal(t, "31/12/1997", student.Birthday)
	assert.Equal(t, time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC), student.CreatedAt)
	assert.Equal(t, time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC), student.UpdatedAt)
}

func TestFindError(t *testing.T) {
	r := mock.StudentRepositoryMock{}
	r.TestError = true

	m := StudentModule{Repository: r}
	id := "25"

	student, apiError := m.Find(id)

	assert.NotNil(t, apiError)
	assert.Equal(t, model.Student{}, student)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestFindNotFoundError(t *testing.T) {
	r := mock.StudentRepositoryMock{}
	r.TestNotFoundError = true

	m := StudentModule{Repository: r}
	id := "25"

	student, apiError := m.Find(id)

	assert.NotNil(t, apiError)
	assert.Equal(t, model.Student{}, student)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.FailMessage, apiError.Message)
	assert.Equal(t, "student not found", apiError.Error)
}

func TestCreate(t *testing.T) {
	r := mock.StudentRepositoryMock{}

	m := StudentModule{Repository: r}
	student := model.Student{
		ID:        "2500651",
		Name:      "Felipe de Azevedo Silva",
		Email:     "felipe9_azevedo@hotmail.com",
		Phone:     "(00)00000-0000",
		Grade:     "7th",
		Birthday:  "31/12/1997",
		CreatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(student)

	assert.Nil(t, apiError)
	assert.NotEqual(t, "", id)
	assert.Equal(t, "2500651", id)
}

func TestCreateError(t *testing.T) {
	r := mock.StudentRepositoryMock{}
	r.TestError = true

	m := StudentModule{Repository: r}
	student := model.Student{
		ID:        "2500651",
		Name:      "Felipe de Azevedo Silva",
		Email:     "felipe9_azevedo@hotmail.com",
		Phone:     "(00)00000-0000",
		Grade:     "7th",
		Birthday:  "31/12/1997",
		CreatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	id, apiError := m.Create(student)

	assert.NotNil(t, apiError)
	assert.Equal(t, "", id)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.CreateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestUpdate(t *testing.T) {
	r := mock.StudentRepositoryMock{}

	m := StudentModule{Repository: r}
	id := "25"
	upStudent := model.Student{
		ID:        "2500651",
		Name:      "Felipe de Azevedo Silva",
		Email:     "felipe9_azevedo@hotmail.com",
		Phone:     "(00)00000-0000",
		Grade:     "7th",
		Birthday:  "31/12/1997",
		CreatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upStudent)

	assert.Nil(t, apiError)
}

func TestUpdateError(t *testing.T) {
	r := mock.StudentRepositoryMock{}
	r.TestUpdateError = true

	m := StudentModule{Repository: r}
	id := "25"
	upStudent := model.Student{
		ID:        "2500651",
		Name:      "Felipe de Azevedo Silva",
		Email:     "felipe9_azevedo@hotmail.com",
		Phone:     "(00)00000-0000",
		Grade:     "7th",
		Birthday:  "31/12/1997",
		CreatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upStudent)

	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestUpdateNotFoundError(t *testing.T) {
	r := mock.StudentRepositoryMock{}
	r.TestNotFoundError = true

	m := StudentModule{Repository: r}
	id := "25"
	upStudent := model.Student{
		ID:        "2500651",
		Name:      "Felipe de Azevedo Silva",
		Email:     "felipe9_azevedo@hotmail.com",
		Phone:     "(00)00000-0000",
		Grade:     "7th",
		Birthday:  "31/12/1997",
		CreatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
		UpdatedAt: time.Date(2021, 04, 05, 04, 00, 00, 00, time.UTC),
	}

	apiError := m.Update(id, upStudent)

	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.UpdateFailMessage, apiError.Message)
	assert.Equal(t, "student not found", apiError.Error)
}

func TestDeleteStudent(t *testing.T) {
	r := mock.StudentRepositoryMock{}

	m := StudentModule{Repository: r}
	id := "25"

	apiError := m.Delete(id)

	assert.Nil(t, apiError)
}

func TestDeleteStudentError(t *testing.T) {
	r := mock.StudentRepositoryMock{}
	r.TestDeleteError = true

	m := StudentModule{Repository: r}
	id := "25"

	apiError := m.Delete(id)

	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusInternalServerError, apiError.Status)
	assert.Equal(t, errors.DeleteFailMessage, apiError.Message)
	assert.Equal(t, "mocked error", apiError.Error)
}

func TestDeleteStudentNotFoundError(t *testing.T) {
	r := mock.StudentRepositoryMock{}
	r.TestNotFoundError = true

	m := StudentModule{Repository: r}
	id := "25"

	apiError := m.Delete(id)

	assert.NotNil(t, apiError)
	assert.Equal(t, http.StatusNotFound, apiError.Status)
	assert.Equal(t, errors.DeleteFailMessage, apiError.Message)
	assert.Equal(t, "student not found", apiError.Error)
}
