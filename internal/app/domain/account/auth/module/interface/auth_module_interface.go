package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/model"
)

type AuthModule interface {
	Login(credentials model.Account) login.Message
	Logout(session model.UserSession) (message login.Message)
}
