package main

import (
	"context"
	"log"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("greet function invoked with %v\n", req)
	return &pb.GreetResponse{
		Result: "Hello " + req.FirstName,
	}, nil
}
