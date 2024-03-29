package grpc

import (
	"context"
	"errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
)

type RegistryGRPCServer struct {
	RegistryModule registries.Module
}

func NewRegistryGRPCServer(module registries.Module) *RegistryGRPCServer {
	return &RegistryGRPCServer{
		RegistryModule: module,
	}
}

func (s *RegistryGRPCServer) Reserve(ctx context.Context, rsv *grpc.ReserveRequest) (*grpc.ReserveResponse, error) {
	registry, err := s.RegistryModule.Find(rsv.GetRegistryNumber())
	if err != nil {
		return &grpc.ReserveResponse{Reserved: false}, errors.New(err.Error)
	}
	switch rsv.Deleted {
	case false:
		registry.Reserved = true
		registry.Available = false
	default:
		registry.Reserved = false
		registry.Available = true
	}
	err = s.RegistryModule.Update(rsv.GetRegistryNumber(), registry)
	if err != nil {
		return &grpc.ReserveResponse{Reserved: false}, errors.New(err.Error)
	}
	return &grpc.ReserveResponse{Reserved: true}, nil
}
