package repository

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserve/model/converter"
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
	result, apiError := r.DB.FetchAllWhere(&[]model.Reserve{}, "book_id", bookId)
	if apiError != nil {
		return nil, apiError
	}
	reserve, apiError := converter.ConvertToSliceReserveObj(result)
	if apiError != nil {
		return nil, apiError
	}
	return reserve, nil
}

func (r ReserveRepository) Find(id string) (model.Reserve, *errors.ApiError) {
	result, apiError := r.DB.Fetch(&model.Reserve{}, id)
	if apiError != nil {
		return model.Reserve{}, apiError
	}
	reserve, apiError := converter.ConvertToReserveObj(result)
	if apiError != nil {
		return model.Reserve{}, apiError
	}
	return reserve, nil
}

func (r ReserveRepository) Create(reserve model.Reserve) (uint, *errors.ApiError) {
	apiError := r.DB.Persist(&reserve)
	if apiError != nil {
		return 0, apiError
	}
	return reserve.ID, nil
}

func (r ReserveRepository) Update(id string, upReserve model.Reserve) *errors.ApiError {
	return r.DB.Refresh(&upReserve, id)
}

func (r ReserveRepository) Delete(id string) *errors.ApiError {
	return r.DB.Remove(&model.Reserve{}, id)
}
