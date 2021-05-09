package pkg

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/book/model"
	categoryModel "github.com/FelipeAz/golibcontrol/internal/app/domain/category/model"
	lendingModel "github.com/FelipeAz/golibcontrol/internal/app/domain/lending/model"
	studentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/student/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAssociateBookInput(t *testing.T) {
	// Init
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{"title":"Code Universe","author":"Unknown Author","registerNumber":"123456"}`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	// Execution
	book, apiError := AssociateBookInput(c)

	// Validation
	assert.Nil(t, apiError)
	assert.Equal(t, "Code Universe", book.Title)
	assert.Equal(t, "Unknown Author", book.Author)
	assert.Equal(t, "123456", book.RegisterNumber)
}

func TestAssociateBookInputWithError(t *testing.T) {
	// Init
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{"title":"Code Universe","author":"Unknown Author","registerNumber":""}`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	// Execution
	book, apiError := AssociateBookInput(c)

	// Validation
	assert.Equal(t, bookModel.Book{}, book)
	assert.Equal(t, http.StatusBadRequest, apiError.Status)
}

func TestAssociateCategoryInput(t *testing.T) {
	// Init
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{"name":"Sci-Fi"}`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	// Execution
	category, apiError := AssociateCategoryInput(c)

	// Validation
	assert.Nil(t, apiError)
	assert.Equal(t, "Sci-Fi", category.Name)
}

func TestAssociateCategoryInputWithError(t *testing.T) {
	// Init
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{"title":""}`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	// Execution
	category, apiError := AssociateCategoryInput(c)

	// Validation
	assert.Equal(t, categoryModel.Category{}, category)
	assert.Equal(t, http.StatusBadRequest, apiError.Status)
}

func TestAssociateStudentInput(t *testing.T) {
	// Init
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(
		`{
				"registerNumber":"123123",
				"name":"Felipe",
				"email":"felipe@hotmail.com",
				"phone":"(00)00000-0000",
				"grade":"7 Ano",
				"birthday":"31/12/1997"
		}`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	// Execution
	student, apiError := AssociateStudentInput(c)

	// Validation
	assert.Nil(t, apiError)
	assert.Equal(t, "123123", student.RegisterNumber)
	assert.Equal(t, "Felipe", student.Name)
	assert.Equal(t, "felipe@hotmail.com", student.Email)
	assert.Equal(t, "(00)00000-0000", student.Phone)
	assert.Equal(t, "7 Ano", student.Grade)
	assert.Equal(t, "31/12/1997", student.Birthday)
}

func TestAssociateStudentInputWithError(t *testing.T) {
	// Init
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(
		`{
				"registerNumber":"",
				"name":"Felipe",
				"email":"felipe@hotmail.com",
				"phone":"(00)00000-0000",
				"grade":"7 Ano",
				"birthday":"31/12/1997"
		}`)

	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	// Execution
	student, apiError := AssociateStudentInput(c)

	// Validation
	assert.Equal(t, studentModel.Student{}, student)
	assert.Equal(t, http.StatusBadRequest, apiError.Status)
}

func TestAssociateLendingInput(t *testing.T) {
	// Init
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{"bookId":10,"studentId":25}`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	// Execution
	lending, apiError := AssociateLendingInput(c)

	// Validation
	assert.Nil(t, apiError)
	assert.Equal(t, 10, int(lending.BookID))
	assert.Equal(t, 25, int(lending.StudentID))
}

func TestAssociateLendingInputWithError(t *testing.T) {
	// Init
	gin.SetMode(gin.TestMode)
	jsonParam := strings.NewReader(`{"book_id":"10"}`)
	req := http.Request{Body: ioutil.NopCloser(jsonParam)}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = &req

	// Execution
	lending, apiError := AssociateLendingInput(c)

	// Validation
	assert.Equal(t, lendingModel.Lending{}, lending)
	assert.Equal(t, http.StatusBadRequest, apiError.Status)
}
