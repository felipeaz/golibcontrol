package service

import (
	"github.com/FelipeAz/golibcontrol/infra/logger"
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

func (s MySQLService) FetchAll(domainObj interface{}) (interface{}, error) {
	result := s.DB.Find(domainObj)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return nil, err
	}
	return domainObj, nil
}

func (s MySQLService) FetchAllWithPreload(domainObj interface{}, preload string) (interface{}, error) {
	result := s.DB.Preload(preload).Find(domainObj)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return nil, err
	}
	return domainObj, nil
}

func (s MySQLService) Fetch(domainObj interface{}, id string) (interface{}, error) {
	result := s.DB.Model(domainObj).Where("id = ?", id).Find(domainObj)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return nil, err
		}
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return domainObj, nil
}

func (s MySQLService) FetchWithPreload(domainObj interface{}, id, preload string) (interface{}, error) {
	result := s.DB.Preload(preload).Model(domainObj).Where("id = ?", id).Find(domainObj)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return nil, err
		}
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return domainObj, nil
}

func (s MySQLService) FetchAllWhere(domainObj interface{}, fieldName, fieldValue string) (interface{}, error) {
	result := s.DB.Model(domainObj).Where(fieldName+" = ? ", fieldValue).Find(domainObj)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return nil, err
		}
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return domainObj, nil
}

func (s MySQLService) FetchAllWhereWithQuery(domainObj interface{}, query string) (interface{}, error) {
	result := s.DB.Model(domainObj).Where(query).Find(domainObj)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			logger.LogError(err)
			return nil, err
		}
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return domainObj, nil
}

func (s MySQLService) Persist(domainObj interface{}) error {
	result := s.DB.Create(domainObj)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (s MySQLService) Refresh(domainObj interface{}, id string) error {
	result := s.DB.Model(domainObj).Where("id = ?", id).Updates(domainObj)
	if err := result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

func (s MySQLService) Remove(domainObj interface{}, id string) error {
	err := s.DB.Where("id = ?", id).Delete(domainObj).Error
	if err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}
