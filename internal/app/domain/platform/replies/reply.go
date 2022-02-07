package replies

import (
	"context"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/grpc"
	grpcServer "github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	grpcService *grpc.Server
)

type Reply struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	CommentId uint        `json:"commentId" gorm:"unique"`
	UserId    uint        `json:"userId" binding:"required"`
	Text      string      `json:"text" binding:"required"`
	Student   interface{} `json:"student" gorm:"-"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}

func (r Reply) TableName() string {
	return "replies"
}

func (r *Reply) AfterFind(tx *gorm.DB) (err error) {
	ctx := context.Background()
	r.Student, err = grpcService.GetStudentInfo(ctx, &grpcServer.GetStudentRequest{
		Id: strconv.Itoa(int(r.UserId)),
	})
	if err != nil {
		return err
	}
	return
}

type Module interface {
	Get(bookId string) ([]Reply, *errors.ApiError)
	Find(id string) (Reply, *errors.ApiError)
	Create(comment Reply) (*Reply, *errors.ApiError)
	Update(id string, upReply Reply) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get(bookId string) ([]Reply, *errors.ApiError)
	Find(id string) (Reply, *errors.ApiError)
	Create(reply Reply) (*Reply, *errors.ApiError)
	Update(id string, upReply Reply) *errors.ApiError
	Delete(id string) *errors.ApiError
}
