package main

import (
	"context"
	"log"

	pb "github.com/Jam3shalliday/grpc-go-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked", c)

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "James",
	})

	if err != nil {
		log.Fatal("Could not greet: %v", err)
	}

	log.Printf("Greeting: $s\n", res.Result)
}
