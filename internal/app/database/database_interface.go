package database

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

type GORMServiceInterface interface {
	Get(domainObj interface{}) (interface{}, *errors.ApiError)
	GetWithPreload(domainObj interface{}, preload string) (interface{}, *errors.ApiError)
	Find(domainObj interface{}, id string) (interface{}, *errors.ApiError)
	FindWithPreload(domainObj interface{}, id, preload string) (interface{}, *errors.ApiError)
	FindWhere(domainObj interface{}, fieldName, fieldValue string) (interface{}, *errors.ApiError)
	FindWhereWithQuery(domainObj interface{}, query string) (interface{}, *errors.ApiError)
	Create(domainObj interface{}) *errors.ApiError
	Update(domainObj interface{}, id string) *errors.ApiError
	Delete(domainObj interface{}, id string) *errors.ApiError
}
