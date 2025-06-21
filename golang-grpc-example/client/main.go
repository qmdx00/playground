package main

import (
	"context"
	"flag"
	pbproto "golang-grpc-example/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:3000", "the address to connect to")
	name = flag.String("name", "world", "Name to greet")
)

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pbproto.NewGreeterClient(conn)

	resp, err := client.SayHello(context.Background(), &pbproto.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("call grpc error: %v", err)
	}

	log.Printf("recv message: %v\n", resp.GetMessage())
}
