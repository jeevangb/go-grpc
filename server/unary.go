package main

import (
	pb "github.com/jeevan/go-grpc-projec/proto"
	"golang.org/x/net/context"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
