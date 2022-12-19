package main

import (
	"context"
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"log"
	"time"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doAverage was invoked")

	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average: %v\n", err)
	}

	log.Printf("Average: %v\n", res.Result)
}
