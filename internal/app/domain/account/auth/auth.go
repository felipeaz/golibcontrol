package auth

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
	"github.com/FelipeAz/golibcontrol/internal/constants/login"
)

type Session struct {
	ConsumerId    string `json:"consumerId" binding:"required"`
	ConsumerKeyId string `json:"consumerKeyId"`
}

func (s Session) TableName() string {
	return "sessions"
}

type Module interface {
	Login(credentials users.Account) login.Data
	Logout(session Session) (message login.Data)
}
