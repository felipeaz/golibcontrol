package repository

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/book/model"
	bookRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/book/repository"
	lendingModel "github.com/FelipeAz/golibcontrol/internal/app/domain/lending/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/lending/model/converter"
	studentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/student/model"
	studentRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/student/repository"
)

// LendingRepository is responsible of getting/saving information from DB.
type LendingRepository struct {
	DB                database.GORMServiceInterface
	StudentRepository studentRepository.StudentRepository
	BookRepository    bookRepository.BookRepository
}

// Get returns all lendings.
func (r LendingRepository) Get() ([]lendingModel.Lending, *errors.ApiError) {
	result, apiError := r.DB.Get(&lendingModel.Lending{})
	if apiError != nil {
		return nil, apiError
	}
	lendings, apiError := converter.ConvertToSliceLendingObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return lendings, nil
}

// Find return one lending from DB by ID.
func (r LendingRepository) Find(id string) (lendingModel.Lending, *errors.ApiError) {
	result, apiError := r.DB.Find(lendingModel.Lending{}, id)
	if apiError != nil {
		return lendingModel.Lending{}, apiError
	}
	lending, apiError := converter.ConvertToLendingObj(result)
	if apiError != nil {
		return lendingModel.Lending{}, apiError
	}
	return lending, nil
}

// Create persist a lending to the DB.
func (r LendingRepository) Create(lending lendingModel.Lending) (uint, *errors.ApiError) {
	validationCH := make(chan *errors.ApiError)
	var wg sync.WaitGroup
	wg.Add(2)

	go func(ch chan<- *errors.ApiError, wg *sync.WaitGroup) {
		apiError := r.BeforeCreateAndUpdate(lending.StudentID, lending.BookID)
		validationCH <- apiError
		wg.Done()
	}(validationCH, &wg)

	go func(ch chan<- *errors.ApiError, wg *sync.WaitGroup) {
		apiError := r.BeforeCreate(lending.StudentID, lending.BookID)
		validationCH <- apiError
		wg.Done()
	}(validationCH, &wg)

	for i := 0; i < 2; i++ {
		err := <-validationCH
		if err != nil {
			err.Message = errors.CreateFailMessage
			return 0, err
		}
	}
	wg.Wait()

	apiError := r.DB.Create(&lending)
	if apiError != nil {
		return 0, apiError
	}

	return lending.ID, nil
}

// Update update an existent lending.
func (r LendingRepository) Update(id string, upLending lendingModel.Lending) (lendingModel.Lending, *errors.ApiError) {
	apiError := r.BeforeCreateAndUpdate(upLending.StudentID, upLending.BookID)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return lendingModel.Lending{}, apiError
	}

	result, apiError := r.DB.Update(&upLending, id)
	if apiError != nil {
		return lendingModel.Lending{}, apiError
	}
	lending, apiError := converter.ConvertToLendingObj(result)
	if apiError != nil {
		return lendingModel.Lending{}, apiError
	}
	return lending, nil
}

// Delete delete an existent lending from DB.
func (r LendingRepository) Delete(id string) *errors.ApiError {
	apiError := r.DB.Delete(&lendingModel.Lending{}, id)
	return apiError
}

// BeforeCreateAndUpdate validate if the student or book exists before create the lending.
func (r LendingRepository) BeforeCreateAndUpdate(studentId, bookId uint) *errors.ApiError {
	_, apiError := r.DB.Find(&studentModel.Student{}, strconv.Itoa(int(studentId)))
	if apiError != nil {
		return apiError
	}

	_, apiError = r.DB.Find(&bookModel.Book{}, strconv.Itoa(int(bookId)))
	if apiError != nil {
		return apiError
	}
	return nil
}

// BeforeCreate validate if the book is already lent.
func (r LendingRepository) BeforeCreate(studentId, bookId uint) *errors.ApiError {
	result, apiError := r.DB.FindWhere(lendingModel.Lending{}, "book_id", strconv.Itoa(int(bookId)))
	if apiError != nil {
		return apiError
	}
	if result != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "book is already lent",
		}
	}

	result, apiError = r.DB.FindWhere(studentModel.Student{}, "student_id", strconv.Itoa(int(studentId)))
	if apiError != nil {
		return apiError
	}
	if result != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   "student has already lent a book",
		}
	}

	return nil
}
