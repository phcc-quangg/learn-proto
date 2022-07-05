package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/Jam3shalliday/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	con, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer con.Close()

	c := pb.NewGreetServiceClient(con)

	doGreet(c)
}
