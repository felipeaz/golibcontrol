package reserves

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/errors"
	"time"
)

type Reserve struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	BookId     uint      `json:"bookId" gorm:"unique"`
	UserId     uint      `json:"userId" gorm:"unique"`
	RetrieveAt time.Time `json:"retrieveAt" time_format:"2006-01-02 15:04:05"`
	CreatedAt  time.Time `time_format:"2006-01-02 15:04:05"`
}

type Module interface {
	Get() ([]Reserve, *errors.ApiError)
	Find(id string) (Reserve, *errors.ApiError)
	Create(comment Reserve) (*Reserve, *errors.ApiError)
	Update(id string, upReserve Reserve) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get() ([]Reserve, *errors.ApiError)
	Find(id string) (Reserve, *errors.ApiError)
	Create(reserve Reserve) (*Reserve, *errors.ApiError)
	Update(id string, upReserve Reserve) *errors.ApiError
	Delete(id string) *errors.ApiError
}
