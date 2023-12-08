package main

import (
	"context"
	"log"
	"time"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var MAX_DEADLINE = 3

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("greetWithDeadline was invoked with", in)

	for i := 0; i < MAX_DEADLINE; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client cancelled the request")
			return nil, status.Error(codes.Canceled, "client cancelled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
