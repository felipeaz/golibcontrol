package grpc

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	grpcserver "github.com/FelipeAz/golibcontrol/internal/app/management/registries/grpc"
	reserve "github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
	"google.golang.org/grpc"
	"net"
)

func Start(registryModule registries.Module) error {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		return err
	}

	grpcSv := grpcserver.NewRegistryGRPCServer(registryModule)
	sv := grpc.NewServer()
	reserve.RegisterReserveServer(sv, grpcSv)

	if err = sv.Serve(listener); err != nil {
		return err
	}
	return nil
}
