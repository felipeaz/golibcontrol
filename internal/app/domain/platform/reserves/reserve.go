package reserves

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

type Reserve struct {
	ID             uint        `json:"id" gorm:"primaryKey"`
	UserId         uint        `json:"userId" gorm:"unique"`
	RegistryNumber int         `json:"registryNumber" gorm:"unique"`
	RetrieveAt     time.Time   `json:"retrieveAt"`
	Book           interface{} `json:"book" gorm:"-"`
	Student        interface{} `json:"student" gorm:"-"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

func (r Reserve) TableName() string {
	return "reserves"
}

func (r *Reserve) BeforeCreate(tx *gorm.DB) error {
	ctx := context.Background()
	resp, err := grpcService.Reserve(ctx, &grpcServer.ReserveRequest{
		RegistryNumber: strconv.Itoa(r.RegistryNumber),
		Deleted:        false,
	})
	if err != nil || resp.Reserved == false {
		return err
	}
	return nil
}

func (r *Reserve) BeforeDelete(tx *gorm.DB) error {
	ctx := context.Background()
	resp, err := grpcService.Reserve(ctx, &grpcServer.ReserveRequest{
		RegistryNumber: strconv.Itoa(r.RegistryNumber),
		Deleted:        true,
	})
	if err != nil || resp.Reserved == false {
		return err
	}
	return nil
}

func (r *Reserve) AfterFind(tx *gorm.DB) (err error) {
	ctx := context.Background()
	r.Book, err = grpcService.GetBookInfo(ctx, &grpcServer.GetBookRequest{
		RegistryNumber: strconv.Itoa(r.RegistryNumber),
	})
	if err != nil {
		return err
	}
	r.Student, err = grpcService.GetStudentInfo(ctx, &grpcServer.GetStudentRequest{
		Id: strconv.Itoa(int(r.UserId)),
	})
	if err != nil {
		return err
	}
	return
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
