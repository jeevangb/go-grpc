package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/jeevan/go-grpc-projec/proto"
)

func HelloBiderctionalStream(client pb.GreetSeviceClient, names *pb.NamesList) {
	log.Printf("Biderctional streaming started")
	stream, err := client.SayHelloBidrectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil { // Parentheses added here
			log.Fatalf("error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional")
}
