package module

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/comment/repository/interface"
)

type CommentModule struct {
	Repository _interface.CommentRepositoryInterface
}
