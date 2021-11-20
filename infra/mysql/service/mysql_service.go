package service

import (
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/logger"
	"gorm.io/gorm"
)

type MySQLService struct {
	DB  *gorm.DB
	Log logger.LogInterface
}

func NewMySQLService(db *gorm.DB, log logger.LogInterface) *MySQLService {
	return &MySQLService{
		DB:  db,
		Log: log,
	}
}

func (s MySQLService) FetchAll(domainObj interface{}) (interface{}, error) {
	result := s.DB.Find(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	return domainObj, nil
}

func (s MySQLService) FetchAllWithPreload(domainObj interface{}, preload string) (interface{}, error) {
	result := s.DB.Preload(preload).Find(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	return domainObj, nil
}

func (s MySQLService) FetchAllWithWhereAndPreload(domainObj interface{}, fieldName, fieldValue, preload string) (interface{}, error) {
	result := s.DB.Preload(preload).Where(fieldName+" = ? ", fieldValue).Find(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	return domainObj, nil
}
func (s MySQLService) FetchAllWithQueryAndPreload(domainObj interface{}, query, preload, join, group string) (interface{}, error) {
	result := s.DB.Joins(join).Preload(preload).Where(query).Group(group).Find(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	return domainObj, nil
}

func (s MySQLService) Fetch(domainObj interface{}, id string) (interface{}, error) {
	result := s.DB.Model(domainObj).Where("id = ?", id).Find(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return domainObj, nil
}

func (s MySQLService) FetchWithPreload(domainObj interface{}, id, preload string) (interface{}, error) {
	result := s.DB.Preload(preload).Model(domainObj).Where("id = ?", id).Find(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return domainObj, nil
}

func (s MySQLService) FetchAllWhere(domainObj interface{}, fieldName, fieldValue string) (interface{}, error) {
	result := s.DB.Model(domainObj).Where(fieldName+" = ? ", fieldValue).Find(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	return domainObj, nil
}

func (s MySQLService) FetchAllWhereWithQuery(domainObj interface{}, query string) (interface{}, error) {
	result := s.DB.Model(domainObj).Where(query).Find(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	return domainObj, nil
}

func (s MySQLService) Persist(domainObj interface{}) error {
	result := s.DB.Create(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}

func (s MySQLService) Refresh(domainObj interface{}, id string) error {
	result := s.DB.Model(domainObj).Where("id = ?", id).Updates(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}

func (s MySQLService) Remove(domainObj interface{}, id string) error {
	err := s.DB.Where("id = ?", id).Delete(domainObj).Error
	if err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}

func (s MySQLService) RemoveWhere(domainObj interface{}, fieldName, fieldValue string) error {
	err := s.DB.Where(fieldName+" = ? ", fieldValue).Delete(domainObj).Error
	if err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}

func (s MySQLService) GetErrorStatusCode(err error) int {
	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
