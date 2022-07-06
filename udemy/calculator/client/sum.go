package main

import (
	"context"
	"log"

	pb "github.com/Jam3shalliday/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("do sum")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 1,
	})

	if err != nil {
		log.Fatal("error", err)
	}

	log.Printf("sum: %d\n", res.Result)
}
