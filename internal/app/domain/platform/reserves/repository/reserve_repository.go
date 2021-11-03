package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves/model/converter"
)

type ReserveRepository struct {
	DB database.GORMServiceInterface
}

func NewReserveRepository(db database.GORMServiceInterface) ReserveRepository {
	return ReserveRepository{
		DB: db,
	}
}

func (r ReserveRepository) Get(bookId string) ([]model.Reserve, *errors.ApiError) {
	result, err := r.DB.FetchAllWhere(&[]model.Reserve{}, "book_id", bookId)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	reserve, apiError := converter.ConvertToSliceReserveObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return reserve, nil
}

func (r ReserveRepository) Find(id string) (model.Reserve, *errors.ApiError) {
	result, err := r.DB.Fetch(&model.Reserve{}, id)
	if err != nil {
		return model.Reserve{}, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	reserve, apiError := converter.ConvertToReserveObj(result)
	if apiError != nil {
		return model.Reserve{}, apiError
	}
	return reserve, nil
}

func (r ReserveRepository) Create(reserve model.Reserve) (uint, *errors.ApiError) {
	err := r.DB.Persist(&reserve)
	if err != nil {
		return 0, &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return reserve.ID, nil
}

func (r ReserveRepository) Update(id string, upReserve model.Reserve) *errors.ApiError {
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
	err := r.DB.Remove(&model.Reserve{}, id)
	if err != nil {
		return &errors.ApiError{
			Status:  r.DB.GetErrorStatusCode(err),
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
