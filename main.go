package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "github.com/jamshoot66/envoy-grpc-example/helloworld"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedExampleServer
}

func (s *server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.EchoResponse{msg: "Hello " + in.GetMsg()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterExampleServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
