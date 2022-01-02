package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/pkg"
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
	result, err := r.DB.FetchAll(&[]reserves.Reserve{})
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
	result, err := r.DB.Fetch(&reserves.Reserve{}, id)
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
	err := r.DB.Persist(&reserve)
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
	err := r.DB.Refresh(&upReserve, id)
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
	err := r.DB.Remove(&reserves.Reserve{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
