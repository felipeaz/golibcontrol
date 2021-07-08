package service

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/infra/logger"
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"gorm.io/gorm"
)

type MySQLService struct {
	DB *gorm.DB
}

func NewMySQLService(db *gorm.DB) (*MySQLService, error) {
	return &MySQLService{
		DB: db,
	}, nil
}

func (s MySQLService) Get(domainObj interface{}) (interface{}, *errors.ApiError) {
	result := s.DB.Find(domainObj)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return domainObj, nil
}

func (s MySQLService) GetWithPreload(domainObj interface{}, preload string) (interface{}, *errors.ApiError) {
	result := s.DB.Preload(preload).Find(domainObj)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailMessage,
			Error:   err.Error(),
		}
	}
	return domainObj, nil
}

func (s MySQLService) Find(domainObj interface{}, id string) (interface{}, *errors.ApiError) {
	result := s.DB.Model(domainObj).Where("id = ?", id).Find(domainObj)
	if result.RowsAffected == 0 {
		return nil, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   errors.ItemNotFoundError,
		}
	}
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.FailMessage,
				Error:   err.Error(),
			}
		}
	}
	return domainObj, nil
}

func (s MySQLService) FindWithPreload(domainObj interface{}, id, preload string) (interface{}, *errors.ApiError) {
	result := s.DB.Preload(preload).Model(domainObj).Where("id = ?", id).Find(domainObj)
	if result.RowsAffected == 0 {
		return nil, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   errors.ItemNotFoundError,
		}
	}
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.FailMessage,
				Error:   err.Error(),
			}
		}
	}
	return domainObj, nil
}

func (s MySQLService) FindWhere(domainObj interface{}, fieldName, fieldValue string) (interface{}, *errors.ApiError) {
	result := s.DB.Model(domainObj).Where(fieldName+" = ? ", fieldValue).Find(domainObj)
	if result.RowsAffected == 0 {
		return nil, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   errors.ItemNotFoundError,
		}
	}
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.FailMessage,
				Error:   err.Error(),
			}
		}
	}
	return domainObj, nil
}

func (s MySQLService) FindWhereWithQuery(domainObj interface{}, query string) (interface{}, *errors.ApiError) {
	result := s.DB.Model(domainObj).Where(query).Find(domainObj)
	if result.RowsAffected == 0 {
		return nil, &errors.ApiError{
			Status:  http.StatusNotFound,
			Message: errors.FailMessage,
			Error:   errors.ItemNotFoundError,
		}
	}
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: errors.FailMessage,
				Error:   err.Error(),
			}
		}
	}
	return domainObj, nil
}

func (s MySQLService) Create(domainObj interface{}) *errors.ApiError {
	result := s.DB.Create(domainObj)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.CreateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (s MySQLService) Update(domainObj interface{}, id string) *errors.ApiError {
	result := s.DB.Model(domainObj).Where("id = ?", id).Updates(domainObj)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.UpdateFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}

func (s MySQLService) Delete(domainObj interface{}, id string) *errors.ApiError {
	err := s.DB.Where("id = ?", id).Delete(domainObj).Error
	if err != nil {
		logger.LogError(err)
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.DeleteFailMessage,
			Error:   err.Error(),
		}
	}
	return nil
}
