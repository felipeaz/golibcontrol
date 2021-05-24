package module

import (
	_interface "github.com/FelipeAz/golibcontrol/internal/app/domain/comment/repository/interface"
)

type CommentModule struct {
	Repository _interface.CommentRepositoryInterface
}
