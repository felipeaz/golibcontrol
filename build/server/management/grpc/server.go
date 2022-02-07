package grpc

import (
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/registries"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	bookGRPC "github.com/FelipeAz/golibcontrol/internal/app/management/books/grpc"
	registryGRPC "github.com/FelipeAz/golibcontrol/internal/app/management/registries/grpc"
	studentGRPC "github.com/FelipeAz/golibcontrol/internal/app/management/students/grpc"
	grpcService "github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
	"google.golang.org/grpc"
	"net"
)

type Modules struct {
	RegistryModule registries.Module
	BookModule     books.Module
	StudentModule  students.Module
}

func Start(modules Modules) error {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		return err
	}

	registryGRPCServer := registryGRPC.NewRegistryGRPCServer(modules.RegistryModule)
	bookGRPCServer := bookGRPC.NewBookGRPCServer(modules.BookModule)
	studentGRPCServer := studentGRPC.NewStudentGRPCServer(modules.StudentModule)
	sv := grpc.NewServer()
	grpcService.RegisterReserveServer(sv, registryGRPCServer)
	grpcService.RegisterGetBookInfoServer(sv, bookGRPCServer)
	grpcService.RegisterGetStudentInfoServer(sv, studentGRPCServer)

	if err = sv.Serve(listener); err != nil {
		return err
	}
	return nil
}
