package usecase

import (
	"context"

	"github.com/jonathangunawan/go-playground/grpc/proto"
)

type Math struct{}

func (ucm Math) Add(ctx context.Context, request *proto.MathRequest) (*proto.MathResponse, error) {
	a, b := request.GetFirstParam(), request.GetSecondParam()

	result := a + b

	return &proto.MathResponse{Result: result}, nil
}

func (ucm Math) Multiply(ctx context.Context, request *proto.MathRequest) (*proto.MathResponse, error) {
	a, b := request.GetFirstParam(), request.GetSecondParam()

	result := a * b

	return &proto.MathResponse{Result: result}, nil
}
