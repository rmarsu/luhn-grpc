package main

import (
	"log"
	"luhn/adder"
	api "luhn/api/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GrpcServer{}
	api.RegisterValidServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
