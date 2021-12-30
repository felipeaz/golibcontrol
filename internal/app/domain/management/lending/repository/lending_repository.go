package repository

import (
	"fmt"
	"strconv"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	lendingModel "github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/model/converter"
)

// LendingRepository is responsible for getting/saving information from DB.
type LendingRepository struct {
	DB database.GORMServiceInterface
}

func NewLendingRepository(db database.GORMServiceInterface) LendingRepository {
	return LendingRepository{
		DB: db,
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
func (r LendingRepository) Create(lending lendingModel.Lending) (*lendingModel.Lending, *errors.ApiError) {
	apiError := r.beforeCreate(lending.StudentID, lending.BookID)
	if apiError != nil {
		return nil, apiError
	}
	err := r.DB.Persist(&lending)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	return &lending, nil
}

// Update update an existent lending.
func (r LendingRepository) Update(id string, upLending lendingModel.Lending) *errors.ApiError {
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

// BeforeCreate validate if the book is already lent.
func (r LendingRepository) beforeCreate(studentId, bookId uint) *errors.ApiError {
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
