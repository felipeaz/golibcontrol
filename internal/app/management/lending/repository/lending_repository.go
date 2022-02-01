package repository

import (
	errorsx "errors"
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/lending"
	"github.com/FelipeAz/golibcontrol/internal/app/filters"
	"github.com/FelipeAz/golibcontrol/internal/app/management/lending/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
	"net/http"
	"strconv"
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
	tx := r.DB.Preload("Book", "Student")
	result, err := r.DB.Find(tx, &[]lending.Lending{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}

	return pkg.ParseToSliceLendingObj(result)
}

// GetByFilter returns all lendings.
func (r LendingRepository) GetByFilter(filter lending.Filter) ([]lending.Lending, *errors.ApiError) {
	queryString := filters.BuildQueryFromFilter(filter)
	tx := r.DB.Preload("Book", "Student")
	tx = r.DB.Where(tx, queryString)
	result, err := r.DB.Find(tx, &[]lending.Lending{})
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
	tx := r.DB.Preload("Book", "Student")
	tx = r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &lending.Lending{})
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
func (r LendingRepository) Create(lendObj lending.Lending) (*lending.Lending, *errors.ApiError) {
	apiError := r.beforeCreate(lendObj.StudentID, lendObj.BookID)
	if apiError != nil {
		return nil, apiError
	}
	err := r.DB.Persist(&lendObj)
	if err != nil {
		if errorsx.Is(err, lending.BookUnavailableError) {
			return nil, &errors.ApiError{
				Status:  http.StatusConflict,
				Message: errors.CreateFailMessage,
				Error:   err.Error(),
			}
		}

		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}

	return &lendObj, nil
}

// Update update an existent lending.
func (r LendingRepository) Update(id string, upLending lending.Lending) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upLending)
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
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &lending.Lending{})
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
	tx := r.DB.Where(nil, fmt.Sprintf("book_id = %s OR student_id = %s",
		strconv.Itoa(int(bookId)),
		strconv.Itoa(int(studentId)),
	))
	result, err := r.DB.Find(tx, &lending.Lending{})
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
