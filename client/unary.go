package main

import (
	"log"
	"time"

	pb "github.com/jeevan/go-grpc-projec/proto"
	"golang.org/x/net/context"
)

func callSayHello(client pb.GreetSeviceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("could not greet:%v", err)
	}
	log.Printf("%s", res.Message)
}
