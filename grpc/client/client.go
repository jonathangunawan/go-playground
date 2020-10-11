package client

import (
	"github.com/jonathangunawan/go-playground/grpc/proto"
	"google.golang.org/grpc"
)

type MathGRPC struct {
	Client proto.MathClient
	conn   *grpc.ClientConn
}

// Make a new wrapper for Dial & New GRPC Client so every gRPC services can define with specific configuration.
func NewMath(address string) (*MathGRPC, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &MathGRPC{
		Client: proto.NewMathClient(conn),
		conn:   conn,
	}, nil
}

func (mg *MathGRPC) Close() {
	mg.conn.Close()
}
