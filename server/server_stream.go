package main

import (
	"log"
	"time"

	pb "github.com/jeevan/go-grpc-projec/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetSevice_SayHelloServerStreamingServer) error {
	log.Printf("got req with names : %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
	return nil
}
