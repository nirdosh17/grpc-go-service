package main

import (
	"log"
	"net"
	"os"

	pb "github.com/nirdosh17/protorepo/greet/lib/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

var addr string = "0.0.0.0:5051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %v: %v\n", addr, err)
	}
	log.Println("listening on", addr)

	opts := []grpc.ServerOption{}
	tls := os.Getenv("TLS")

	if tls == "true" {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalln("failed loading credentials", err)
		}
		log.Println("TLS enabled")
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	// we need to register this server
	pb.RegisterGreetServiceServer(s, &Server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err = s.Serve(lis); err != nil {
		log.Fatalln("failed to serve:", err)
	}
}
