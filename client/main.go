package main

import (
	"log"

	pb "github.com/jeevan/go-grpc-projec/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("not connect :%v", err)
	}
	defer conn.Close()

	client := pb.NewGreetSeviceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Appu", "Jeevan", "Nithin"},
	}
	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	callSayHelloClientStream(client, names)
	// callHelloBiderctionalStream(client, names)
}
