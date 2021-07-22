package repository

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	bookModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/model"
	bookRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/management/book/repository/interface"
	lendingModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model/converter"
	studentModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/model"
	studentRepository "github.com/FelipeAz/golibcontrol/internal/app/domain/management/student/repository/interface"
)

// LendingRepository is responsible of getting/saving information from DB.
type LendingRepository struct {
	DB                database.GORMServiceInterface
	StudentRepository studentRepository.StudentRepositoryInterface
	BookRepository    bookRepository.BookRepositoryInterface
}

func NewLendingRepository(db database.GORMServiceInterface, stRepo studentRepository.StudentRepositoryInterface, bRepo bookRepository.BookRepositoryInterface) LendingRepository {
	return LendingRepository{
		DB:                db,
		StudentRepository: stRepo,
		BookRepository:    bRepo,
	}
}

// Get returns all lendings.
func (r LendingRepository) Get() ([]lendingModel.Lending, *errors.ApiError) {
	result, apiError := r.DB.Get(&[]lendingModel.Lending{})
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
	result, apiError := r.DB.Find(&lendingModel.Lending{}, id)
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
func (r LendingRepository) Update(id string, upLending lendingModel.Lending) *errors.ApiError {
	apiError := r.BeforeCreateAndUpdate(upLending.StudentID, upLending.BookID)
	if apiError != nil {
		apiError.Message = errors.UpdateFailMessage
		return apiError
	}
	return r.DB.Update(&upLending, id)
}

// Delete delete an existent lending from DB.
func (r LendingRepository) Delete(id string) *errors.ApiError {
	return r.DB.Delete(&lendingModel.Lending{}, id)
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
	result, apiError := r.DB.FindWhereWithQuery(
		&lendingModel.Lending{},
		fmt.Sprintf("book_id = %s OR student_id = %s", strconv.Itoa(int(bookId)), strconv.Itoa(int(studentId))))
	if apiError != nil && apiError.Error != errors.ItemNotFoundError {
		return apiError
	}
	if result != nil && apiError == nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   errors.LendingNotAvailableError,
		}
	}
	return nil
}
