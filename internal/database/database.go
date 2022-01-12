package database

import "gorm.io/gorm"

type GORMServiceInterface interface {
	Find(tx *gorm.DB, domainObj interface{}) (interface{}, error)
	FindOne(tx *gorm.DB, domainObj interface{}) (interface{}, error)
	Persist(domainObj interface{}) error
	Refresh(tx *gorm.DB, domainObj interface{}) error
	Set(tx *gorm.DB, domainObj interface{}, field string, value interface{}) error
	Remove(tx *gorm.DB, domainObj interface{}) error
	Preload(preload ...string) *gorm.DB
	Join(tx *gorm.DB, join ...string) *gorm.DB
	Group(tx *gorm.DB, group ...string) *gorm.DB
	Where(tx *gorm.DB, where string) *gorm.DB
	GetErrorStatusCode(err error) int
}
