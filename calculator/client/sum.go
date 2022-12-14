package main

import (
	"context"
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	r, err := c.Sum(context.Background(), &pb.CalculateRequest{
		FirstNumber:  10,
		SecondNumber: 5,
	})

	if err != nil {
		log.Fatalf("Could not doSum: %v\n", err)
	}

	log.Printf("Result of summatory: %d\n", r.Result)
}
