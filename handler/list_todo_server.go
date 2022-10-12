package handler

import (
	"context"
	"list-todo-ms/grpc/generated"

	"google.golang.org/protobuf/types/known/emptypb"
)

type handler struct {
	generated.UnimplementedListToDoServer
}

func NewHandler() *handler {
	return &handler{}
}

func (handler) HealthCheck(context.Context, *emptypb.Empty) (*generated.HealthCheckResponse, error) {
	return &generated.HealthCheckResponse{
		Result: "service running",
	}, nil
}
