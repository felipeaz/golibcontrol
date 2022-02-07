package grpc

import (
	"context"
	"errors"
	"github.com/FelipeAz/golibcontrol/internal/app/domain/management/books"
	"github.com/FelipeAz/golibcontrol/internal/app/plugins/grpc"
)

var (
	BookNotFound = errors.New("book not found")
)

type BookGRPCServer struct {
	BookModule books.Module
}

func NewBookGRPCServer(module books.Module) *BookGRPCServer {
	return &BookGRPCServer{
		BookModule: module,
	}
}

func (s *BookGRPCServer) GetBookInfo(ctx context.Context, req *grpc.GetBookRequest) (*grpc.GetBookResponse, error) {
	var book books.Book

	switch {
	case req.Id != "":
		resp, apiErr := s.BookModule.Find(req.GetId())
		if apiErr != nil {
			return nil, errors.New(apiErr.Error)
		}
		book = resp
	default:
		resp, err := s.BookModule.GetByFilter(books.Filter{
			RegistryNumber: req.GetRegistryNumber(),
		})
		if err != nil {
			return nil, errors.New(err.Error)
		}
		if len(resp) == 0 {
			return nil, BookNotFound
		}
		book = resp[0]
	}

	return &grpc.GetBookResponse{
		Title:  book.Title,
		Author: book.Author,
		Image:  book.Image,
	}, nil
}
