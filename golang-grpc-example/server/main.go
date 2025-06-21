package main

import (
	"context"
	"flag"
	"fmt"
	pbproto "golang-grpc-example/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreetServer struct {
	pbproto.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *GreetServer) SayHello(ctx context.Context, in *pbproto.HelloRequest) (*pbproto.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pbproto.HelloReply{Message: "Hello " + in.GetName()}, nil
}

var (
	port = flag.Int("port", 3000, "The server port")
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	s := grpc.NewServer()

	pbproto.RegisterGreeterServer(s, &GreetServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
