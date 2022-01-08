package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/groups"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type GroupModule struct {
	Repository groups.GroupRepositoryInterface
	Log        logger.LogInterface
}

func NewGroupModule(repo groups.GroupRepositoryInterface, log logger.LogInterface) GroupModule {
	return GroupModule{
		Repository: repo,
		Log:        log,
	}
}

func (m GroupModule) Get() ([]groups.Group, *errors.ApiError) {
	return m.Repository.Get()
}

func (m GroupModule) Find(id string) (groups.Group, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m GroupModule) Create(comment groups.Group) (*groups.Group, *errors.ApiError) {
	return m.Repository.Create(comment)
}

func (m GroupModule) Update(id string, upGroup groups.Group) *errors.ApiError {
	return m.Repository.Update(id, upGroup)
}

func (m GroupModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
