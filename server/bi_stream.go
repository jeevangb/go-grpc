package main

import (
	"io"
	"log"

	pb "github.com/jeevan/go-grpc-projec/proto"
)

func (s *helloServer) SayHelloBidrectionalStreaming(stream pb.GreetSevice_SayHelloBidrectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("got req with  name: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
