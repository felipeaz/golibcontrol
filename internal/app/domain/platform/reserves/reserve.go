package reserves

import (
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"time"
)

type Reserve struct {
	ID         uint      `json:"id" gorm:"primarykey"`
	BookId     uint      `json:"bookId" gorm:"unique"`
	UserId     uint      `json:"userId" gorm:"unique"`
	RetrieveAt time.Time `json:"retrieveAt"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (r Reserve) TableName() string {
	return "reserves"
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
