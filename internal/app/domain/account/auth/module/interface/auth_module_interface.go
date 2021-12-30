package _interface

import (
	"github.com/FelipeAz/golibcontrol/internal/app/constants/login"
	session_model "github.com/FelipeAz/golibcontrol/internal/app/domain/account/auth/model"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/account/users/model"
)

type AuthModule interface {
	Login(credentials model.Account) login.Message
	Logout(session session_model.UserSession) (message login.Message)
}
