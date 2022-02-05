package repository

import (
	"fmt"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	"github.com/FelipeAz/golibcontrol/internal/app/platform/reserves/pkg"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/database"
	"gorm.io/gorm"
)

type ReserveRepository struct {
	DB database.GORMServiceInterface
}

func NewReserveRepository(db database.GORMServiceInterface) ReserveRepository {
	return ReserveRepository{
		DB: db,
	}
}

func (r ReserveRepository) Get() ([]reserves.Reserve, *errors.ApiError) {
	result, err := r.DB.Find(nil, &[]reserves.Reserve{})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToSliceReserveObj(result)
}

func (r ReserveRepository) Find(id string) (reserves.Reserve, *errors.ApiError) {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	result, err := r.DB.FindOne(tx, &reserves.Reserve{})
	if err != nil {
		return reserves.Reserve{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return pkg.ParseToReserveObj(result)
}

func (r ReserveRepository) Create(reserve reserves.Reserve) (*reserves.Reserve, *errors.ApiError) {
	tx := r.DB.GetTx()
	err := tx.Transaction(func(tx *gorm.DB) error {
		txErr := r.DB.Persist(&reserve)
		if txErr != nil {
			return txErr
		}
		return nil
	})
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return &reserve, nil
}

func (r ReserveRepository) Update(id string, upReserve reserves.Reserve) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Refresh(tx, &upReserve)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (r ReserveRepository) Delete(id string) *errors.ApiError {
	tx := r.DB.Where(nil, fmt.Sprintf("id = %s", id))
	err := r.DB.Remove(tx, &reserves.Reserve{})
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
