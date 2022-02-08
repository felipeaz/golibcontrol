package comments

import (
	"context"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/grpc"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/platform/replies"
	grpcServer "github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	grpcService *grpc.Server
)

type Comment struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	BookId    uint            `json:"bookId" binding:"required"`
	UserId    uint            `json:"userId" binding:"required"`
	Text      string          `json:"text" binding:"required"`
	Reply     []replies.Reply `json:"replies"`
	Book      interface{}     `json:"book" gorm:"-"`
	Student   interface{}     `json:"student" gorm:"-"`
	CreatedAt time.Time       `json:"createdAt"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

func (c Comment) TableName() string {
	return "comments"
}

func (c *Comment) AfterFind(tx *gorm.DB) (err error) {
	ctx := context.Background()
	c.Book, err = grpcService.GetBookInfo(ctx, &grpcServer.GetBookRequest{
		Id: strconv.Itoa(int(c.BookId)),
	})
	if err != nil {
		return err
	}
	c.Student, err = grpcService.GetStudentInfo(ctx, &grpcServer.GetStudentRequest{
		Id: strconv.Itoa(int(c.UserId)),
	})
	if err != nil {
		return err
	}
	return
}

type Module interface {
	Get(bookId string) ([]Comment, *errors.ApiError)
	Find(id string) (Comment, *errors.ApiError)
	Create(comment Comment) (*Comment, *errors.ApiError)
	Update(id string, upComment Comment) *errors.ApiError
	Delete(id string) *errors.ApiError
}

type Repository interface {
	Get(bookId string) ([]Comment, *errors.ApiError)
	Find(id string) (Comment, *errors.ApiError)
	Create(comment Comment) (*Comment, *errors.ApiError)
	Update(id string, upComment Comment) *errors.ApiError
	Delete(id string) *errors.ApiError
}
