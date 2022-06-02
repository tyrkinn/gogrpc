package main

import (
	"context"
	"gogrpc/proto"
	"gogrpc/utils"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	proto.UnimplementedAdderServer
}

func (s *server) Add(ctx context.Context, in *proto.AddRequest) (*proto.AddResponse, error) {
	log.Printf("Calculating %d + %d", in.X, in.Y)
	return &proto.AddResponse{Result: in.X + in.Y}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	utils.CheckErr(err, nil)
	s := grpc.NewServer()
	proto.RegisterAdderServer(s, &server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
