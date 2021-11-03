package repository

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/model"
	bookRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/management/books/repository/interface"
	lendingModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model/converter"
	studentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/students/model"
	studentRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/management/students/repository/interface"
)

// LendingRepository is responsible for getting/saving information from DB.
type LendingRepository struct {
	DB                database.GORMServiceInterface
	StudentRepository studentRepository.StudentRepositoryInterface
	BookRepository    bookRepository.BookRepositoryInterface
}

func NewLendingRepository(
	db database.GORMServiceInterface,
	stRepo studentRepository.StudentRepositoryInterface,
	bRepo bookRepository.BookRepositoryInterface,
) LendingRepository {
	return LendingRepository{
		DB:                db,
		StudentRepository: stRepo,
		BookRepository:    bRepo,
	}
}

// Get returns all lendings.
func (r LendingRepository) Get() ([]lendingModel.Lending, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]lendingModel.Lending{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}

	lendings, apiError := converter.ConvertToSliceLendingObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return lendings, nil
}

// Find return one lending from DB by ID.
func (r LendingRepository) Find(id string) (lendingModel.Lending, *errors.ApiError) {
	result, err := r.DB.Fetch(&lendingModel.Lending{}, id)
	if err != nil {
		return lendingModel.Lending{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
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

	err := r.DB.Persist(&lending)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	return lending.ID, nil
}

// Update update an existent lending.
func (r LendingRepository) Update(id string, upLending lendingModel.Lending) *errors.ApiError {
	apiError := r.BeforeCreateAndUpdate(upLending.StudentID, upLending.BookID)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return apiError
	}
	err := r.DB.Refresh(&upLending, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// Delete delete an existent lending from DB.
func (r LendingRepository) Delete(id string) *errors.ApiError {
	err := r.DB.Remove(&lendingModel.Lending{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

// BeforeCreateAndUpdate validate if the student or book exists before create the lending.
func (r LendingRepository) BeforeCreateAndUpdate(studentId, bookId uint) *errors.ApiError {
	student, err := r.DB.Fetch(&studentModel.Student{}, strconv.Itoa(int(studentId)))
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	if student == nil {
		return &errors.ApiError{
			Status: r.DB.GetErrorStatusCode(err),
			Error:  errors.StudentNotFoundError,
		}
	}

	book, err := r.DB.Fetch(&bookModel.Book{}, strconv.Itoa(int(bookId)))
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	if book == nil {
		return &errors.ApiError{
			Status: r.DB.GetErrorStatusCode(err),
			Error:  errors.BookNotFoundError,
		}
	}
	return nil
}

// BeforeCreate validate if the book is already lent.
func (r LendingRepository) BeforeCreate(studentId, bookId uint) *errors.ApiError {
	result, err := r.DB.FetchAllWhereWithQuery(
		&lendingModel.Lending{},
		fmt.Sprintf("book_id = %s OR student_id = %s", strconv.Itoa(int(bookId)), strconv.Itoa(int(studentId))))
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	if result != nil {
		return &errors.ApiError{
			Status: r.DB.GetErrorStatusCode(err),
			Error:  errors.LendingNotAvailableError,
		}
	}
	return nil
}
