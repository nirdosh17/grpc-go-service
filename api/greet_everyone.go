package main

import (
	"io"
	"log"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone invoked!")
	// infinite loop to receive something from client
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalln("error while reading client stream", err)
		}

		err = stream.Send(&pb.GreetResponse{
			Result: "Hello " + req.FirstName + "!",
		})

		if err != nil {
			log.Fatalln("error sending data to client", err)
		}
	}
}
