package grpc_server

import (
	"google.golang.org/grpc"
)

func InitSignGrpcConn(host string) (*grpc.ClientConn, error) {
	return grpc.Dial(host, grpc.WithInsecure())
}
