package main

import (
	"context"
	"log"

	pb "github.com/Jam3shalliday/sum"
)

func doSum(c pb.SumServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		number_1: 12,
		number_2: 13,
	})

	if err != nil {
		log.Fatal("Could not greet: %v", err)
	}

	log.Printf("Greeting: $s\n", res.Result)
}
