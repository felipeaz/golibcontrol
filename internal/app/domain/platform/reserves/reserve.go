package reserves

import (
	"context"
	reserve "github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
	"github.com/FelipeAz/golibcontrol/internal/constants/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var (
	grpcAddr = "management-service:4040"
)

type Reserve struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserId         uint      `json:"userId" gorm:"unique"`
	RegistryNumber int       `json:"registryNumber" gorm:"unique"`
	RetrieveAt     time.Time `json:"retrieveAt"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

func (r Reserve) TableName() string {
	return "reserves"
}

func (r *Reserve) BeforeCreate(tx *gorm.DB) error {
	ctx := context.Background()
	cli, err := connectToGRPCClient()
	resp, err := cli.Reserve(ctx, &reserve.ReserveRequest{
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
	cli, err := connectToGRPCClient()
	resp, err := cli.Reserve(ctx, &reserve.ReserveRequest{
		RegistryNumber: strconv.Itoa(r.RegistryNumber),
		Deleted:        true,
	})
	if err != nil || resp.Reserved == false {
		return err
	}
	return nil
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

func connectToGRPCClient() (reserve.ReserveClient, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.Dial(grpcAddr, opts)
	if err != nil {
		return nil, err
	}
	return reserve.NewReserveClient(cc), nil
}
