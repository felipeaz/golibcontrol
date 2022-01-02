package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/consumer"
	domain "github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"net/http"

	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	databaseInterface "github.com/FelipeAz/golibcontrol/internal/app/database"
	"github.com/FelipeAz/golibcontrol/internal/app/logger"
)

type AccountModule struct {
	Repository domain.Repository
	Consumer   consumer.Interface
	Cache      databaseInterface.CacheInterface
	Log        logger.LogInterface
}

func NewAccountModule(
	repo domain.Repository,
	consumer consumer.Interface,
	cache databaseInterface.CacheInterface,
	log logger.LogInterface,
) AccountModule {
	return AccountModule{
		Repository: repo,
		Consumer:   consumer,
		Cache:      cache,
		Log:        log,
	}
}

// Get returns all accounts.
func (m AccountModule) Get() ([]domain.Account, *errors.ApiError) {
	return m.Repository.Get()
}

// Find return one user by ID.
func (m AccountModule) Find(id string) (domain.Account, *errors.ApiError) {
	return m.Repository.Find(id)
}

// Create creates a user
func (m AccountModule) Create(account domain.Account) (*domain.Account, *errors.ApiError) {
	cons, err := m.Consumer.CreateConsumer(account.Email)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToCreateConsumer,
			Error:   err.Error(),
		}
	}
	account.ConsumerId = cons.Id

	return m.Repository.Create(account)
}

// Update update an existent user.
func (m AccountModule) Update(id string, upAccount domain.Account) *errors.ApiError {
	return m.Repository.Update(id, upAccount)
}

// Delete delete an existent user by id.
func (m AccountModule) Delete(id string) *errors.ApiError {
	user, apiError := m.Repository.Find(id)
	if apiError != nil {
		return apiError
	}

	err := m.Consumer.DeleteConsumer(user.ConsumerId)
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToDeleteConsumer,
			Error:   err.Error(),
		}
	}

	err = m.Cache.Flush(user.ConsumerId)
	if err != nil {
		return &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: errors.FailedToDeleteAuthenticationOnCache,
			Error:   err.Error(),
		}
	}

	return m.Repository.Delete(id)
}
