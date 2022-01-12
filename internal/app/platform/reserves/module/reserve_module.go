package module

import (
	"encoding/json"
	kafkaErrors "github.com/FelipeAz/golibcontrol/infra/kafka/errors"
	"github.com/FelipeAz/golibcontrol/infra/kafka/producer"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/reserves"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"github.com/FelipeAz/golibcontrol/internal/logger"
	"net/http"
)

const (
	ReserveTopic = "reserves"
)

type ReserveModule struct {
	Repository reserves.Repository
	Log        logger.LogInterface
	Producer   producer.ProducerInterface
}

func NewReserveModule(repo reserves.Repository, log logger.LogInterface) ReserveModule {
	return ReserveModule{
		Repository: repo,
		Log:        log,
	}
}

func (m ReserveModule) Get() ([]reserves.Reserve, *errors.ApiError) {
	return m.Repository.Get()
}

func (m ReserveModule) Find(id string) (reserves.Reserve, *errors.ApiError) {
	return m.Repository.Find(id)
}

func (m ReserveModule) Create(reserve reserves.Reserve) (*reserves.Reserve, *errors.ApiError) {
	b, err := json.Marshal(reserve)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: errors.FailedToMarshal,
			Error:   err.Error(),
		}
	}
	err = m.Producer.Produce(ReserveTopic, b)
	if err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusBadRequest,
			Message: kafkaErrors.ProduceFailed,
			Error:   err.Error(),
		}
	}
	return m.Repository.Create(reserve)
}

func (m ReserveModule) Update(id string, upReserve reserves.Reserve) *errors.ApiError {
	return m.Repository.Update(id, upReserve)
}

func (m ReserveModule) Delete(id string) *errors.ApiError {
	return m.Repository.Delete(id)
}
