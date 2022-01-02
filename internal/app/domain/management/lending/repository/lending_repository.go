package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending/pkg"
	"strconv"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
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
func (r LendingRepository) Get() ([]lending.Lending, *errors.ApiError) {
	result, err := r.DB.FetchAll(&[]lending.Lending{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}

	return pkg.ParseToSliceLendingObj(result)
}

// Find return one lending from DB by ID.
func (r LendingRepository) Find(id string) (lending.Lending, *errors.ApiError) {
	result, err := r.DB.Fetch(&lending.Lending{}, id)
	if err != nil {
		return lending.Lending{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToLendingObj(result)
}

// Create persist a lending to the DB.
func (r LendingRepository) Create(lending lending.Lending) (*lending.Lending, *errors.ApiError) {
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
func (r LendingRepository) Update(id string, upLending lending.Lending) *errors.ApiError {
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
	err := r.DB.Remove(&lending.Lending{}, id)
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
		&lending.Lending{},
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
