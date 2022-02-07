package grpc

import (
	"context"
	grpcServer "github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

var (
	managementGRPCServerAddr = os.Getenv("MANAGEMENT_GRPC_SERVER_HOST")
)

type Server struct{}

func (s *Server) Reserve(ctx context.Context, req *grpcServer.ReserveRequest) (*grpcServer.ReserveResponse, error) {
	cli, err := s.connectToReserveClient()
	if err != nil {
		return nil, err
	}
	return cli.Reserve(ctx, req)
}

func (s *Server) GetBookInfo(ctx context.Context, req *grpcServer.GetBookRequest) (*grpcServer.GetBookResponse, error) {
	cli, err := s.connectToBookClient()
	if err != nil {
		return nil, err
	}
	return cli.GetBookInfo(ctx, req)
}

func (s *Server) GetStudentInfo(ctx context.Context, req *grpcServer.GetStudentRequest) (*grpcServer.GetStudentResponse, error) {
	cli, err := s.connectToStudentClient()
	if err != nil {
		return nil, err
	}
	return cli.GetStudentInfo(ctx, req)
}

func (s *Server) getCli(addr string) (*grpc.ClientConn, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	return grpc.Dial(addr, opts)
}

func (s *Server) connectToReserveClient() (grpcServer.ReserveClient, error) {
	cc, err := s.getCli(managementGRPCServerAddr)
	if err != nil {
		return nil, err
	}
	return grpcServer.NewReserveClient(cc), nil
}

func (s *Server) connectToStudentClient() (grpcServer.GetStudentInfoClient, error) {
	cc, err := s.getCli(managementGRPCServerAddr)
	if err != nil {
		return nil, err
	}
	return grpcServer.NewGetStudentInfoClient(cc), nil
}

func (s *Server) connectToBookClient() (grpcServer.GetBookInfoClient, error) {
	cc, err := s.getCli(managementGRPCServerAddr)
	if err != nil {
		return nil, err
	}
	return grpcServer.NewGetBookInfoClient(cc), nil
}
