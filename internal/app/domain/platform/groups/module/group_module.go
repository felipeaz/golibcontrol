package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups/model"
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups/repository/interface"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type GroupModule struct {
	Repository _interface.GroupRepositoryInterface
	Log        logger.LogInterface
}

func NewGroupModule(repo _interface.GroupRepositoryInterface, log logger.LogInterface) GroupModule {
	return GroupModule{
		Repository: repo,
		Log:        log,
	}
}

func (m GroupModule) Get() ([]model.Group, *errors.ApiError) {
	return m.Repository.Get()
}

func (m GroupModule) Find(id string) (model.Group, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m GroupModule) Create(comment model.Group) (uint, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m GroupModule) Update(id string, upGroup model.Group) *errors.ApiError {
	return m.Repository.Update(id, upGroup)
}

func (m GroupModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
