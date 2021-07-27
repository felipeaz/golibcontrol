package database

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
)

type GORMServiceInterface interface {
	FetchAll(domainObj interface{}) (interface{}, *errors.ApiError)
	FetchAllWithPreload(domainObj interface{}, preload string) (interface{}, *errors.ApiError)
	Fetch(domainObj interface{}, id string) (interface{}, *errors.ApiError)
	FetchWithPreload(domainObj interface{}, id, preload string) (interface{}, *errors.ApiError)
	FetchAllWhere(domainObj interface{}, fieldName, fieldValue string) (interface{}, *errors.ApiError)
	FetchAllWhereWithQuery(domainObj interface{}, query string) (interface{}, *errors.ApiError)
	Persist(domainObj interface{}) *errors.ApiError
	Refresh(domainObj interface{}, id string) *errors.ApiError
	Remove(domainObj interface{}, id string) *errors.ApiError
}
