package service

import (
	"errors"
	"github.com/FelipeAz/golibcontrol/internal/logger"
	"net/http"

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

func (s MySQLService) Find(tx *gorm.DB, domainObj interface{}) (interface{}, error) {
	var result *gorm.DB
	if tx != nil {
		result = tx.Find(domainObj)
	} else {
		result = s.DB.Find(domainObj)
	}
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return domainObj, nil
}

func (s MySQLService) FindOne(tx *gorm.DB, domainObj interface{}) (interface{}, error) {
	result := tx.First(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
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

func (s MySQLService) Refresh(tx *gorm.DB, domainObj interface{}) error {
	result := tx.Updates(domainObj)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}

func (s MySQLService) Set(tx *gorm.DB, domainObj interface{}, fieldName string, fieldValue interface{}) error {
	result := tx.Model(domainObj).Update(fieldName, fieldValue)
	if err := result.Error; err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}

func (s MySQLService) Remove(tx *gorm.DB, domainObj interface{}) error {
	err := tx.Delete(domainObj).Error
	if err != nil {
		s.Log.Error(err)
		return err
	}
	return nil
}

func (s MySQLService) Preload(preload ...string) *gorm.DB {
	if preload == nil {
		return s.DB
	}
	tx := s.DB
	for _, v := range preload {
		tx = tx.Preload(v)
	}
	return tx
}

func (s MySQLService) Join(tx *gorm.DB, join ...string) *gorm.DB {
	if join == nil {
		return tx
	}
	txx := tx
	for _, v := range join {
		txx = tx.Joins(v)
	}
	return txx
}

func (s MySQLService) Group(tx *gorm.DB, group ...string) *gorm.DB {
	if group == nil {
		return tx
	}
	txx := tx
	for _, v := range group {
		txx = tx.Group(v)
	}
	return txx
}

func (s MySQLService) Where(tx *gorm.DB, where string) *gorm.DB {
	if where == "" {
		return tx
	}
	if tx == nil {
		return s.DB.Where(where)
	}
	return tx.Where(where)
}

func (s MySQLService) GetTx() *gorm.DB {
	return s.DB
}

func (s MySQLService) GetErrorStatusCode(err error) int {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
