package main

import (
	"context"
	"log"
	"net"

	"github.com/ntnbrtnkv/otus-golang/calendar/model"
	grpc_gen "github.com/ntnbrtnkv/otus-golang/calendar/proto"
	"github.com/ntnbrtnkv/otus-golang/calendar/storage/inmemory"
	grpc "google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	grpc_gen.UnimplementedCalendarServer
	storage inmemory.InMemoryStorage
}

func (s *Server) ListEvents(context.Context, *emptypb.Empty) (*grpc_gen.Events, error) {
	events := s.storage.ListEvents()
	var resp []*grpc_gen.Event
	for _, event := range events {
		resp = append(resp, &grpc_gen.Event{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			TimeStart:   timestamppb.New(event.TimeStart),
			TimeEnd:     timestamppb.New(event.TimeEnd),
		})
	}

	return &grpc_gen.Events{
		Events: resp,
	}, nil
}

func (s *Server) ChangeEvent(ctx context.Context, event *grpc_gen.Event) (*grpc_gen.Result, error) {
	modelEvent := model.Event{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		TimeStart:   event.TimeStart.AsTime(),
		TimeEnd:     event.TimeEnd.AsTime(),
	}

	err := s.storage.ChangeEvent(modelEvent)

	if err != nil {
		return &grpc_gen.Result{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}

	return &grpc_gen.Result{
		Ok: true,
	}, nil
}

func (s *Server) AddEvent(cx context.Context, event *grpc_gen.Event) (*grpc_gen.Result, error) {
	err := s.storage.AddEvent(model.Event{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		TimeStart:   event.TimeStart.AsTime(),
		TimeEnd:     event.TimeEnd.AsTime(),
	})

	if err != nil {
		return &grpc_gen.Result{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}

	return &grpc_gen.Result{
		Ok: true,
	}, nil
}

func (s *Server) RemoveEvent(cx context.Context, event *grpc_gen.IDEventMessage) (*grpc_gen.Result, error) {
	err := s.storage.RemoveEvent(event.ID)

	if err != nil {
		return &grpc_gen.Result{
			Ok:    false,
			Error: err.Error(),
		}, nil
	}

	return &grpc_gen.Result{
		Ok: true,
	}, nil
}

func NewGRPCServer() *grpc.Server {
	storage := inmemory.InMemoryStorage{}
	gsrv := grpc.NewServer()

	srv := &Server{
		storage: storage,
	}

	grpc_gen.RegisterCalendarServer(gsrv, srv)

	return gsrv
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8093")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := NewGRPCServer()
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
