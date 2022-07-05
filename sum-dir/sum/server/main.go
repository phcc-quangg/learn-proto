package main

import (
	"log"
	"net"

	pb "github.com/Jam3shalliday/sum/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:8080"

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal("Failed to listen: %v\n", addr)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
