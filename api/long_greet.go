package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		// this means client will not send request anymore
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}
		log.Println("LongGreet message received", req.FirstName)

		if err != nil {
			log.Fatalln("error while reading client stream", err)
		}

		res += fmt.Sprintf(" Hello %s! ", req.FirstName)
	}
}
