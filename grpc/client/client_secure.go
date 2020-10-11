package client

import (
	"context"
	"fmt"

	"github.com/jonathangunawan/go-playground/grpc/proto"
	"google.golang.org/grpc"
)

type MathGRPCSecure struct {
	client proto.MathClient
	conn   *grpc.ClientConn
}

func NewMathSecure(address string) (*MathGRPCSecure, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &MathGRPCSecure{
		client: proto.NewMathClient(conn),
		conn:   conn,
	}, nil
}

func (c *MathGRPCSecure) Close() {
	c.conn.Close()
}

func (mgs *MathGRPCSecure) Add(ctx context.Context, req *proto.MathRequest) (*proto.MathResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("no request")
	}

	return mgs.client.Add(ctx, req)
}

func (mgs *MathGRPCSecure) Multiply(ctx context.Context, req *proto.MathRequest) (*proto.MathResponse, error) {
	if req == nil {
		return nil, fmt.Errorf("no request")
	}

	return mgs.client.Multiply(ctx, req)
}
