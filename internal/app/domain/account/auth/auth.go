package auth

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users"
)

type Session struct {
	ConsumerId    string `json:"consumerId" binding:"required"`
	ConsumerKeyId string `json:"consumerKeyId"`
}

func (s Session) TableName() string {
	return "sessions"
}

type Module interface {
	Login(credentials users.Account) login.Message
	Logout(session Session) (message login.Message)
}
