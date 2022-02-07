package grpc

import (
	"context"
	"errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/students"
	"github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
)

type StudentGRPCServer struct {
	StudentModule students.Module
}

func NewStudentGRPCServer(module students.Module) *StudentGRPCServer {
	return &StudentGRPCServer{
		StudentModule: module,
	}
}

func (s *StudentGRPCServer) GetStudentInfo(ctx context.Context, req *grpc.GetStudentRequest) (*grpc.GetStudentResponse, error) {
	student, err := s.StudentModule.Find(req.GetId())
	if err != nil {
		return nil, errors.New(err.Error)
	}
	return &grpc.GetStudentResponse{
		Name:  student.Name,
		Email: student.Email,
		Phone: student.Phone,
	}, nil
}
