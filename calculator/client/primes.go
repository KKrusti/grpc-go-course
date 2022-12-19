package main

import (
	"context"
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func doPrime(c pb.CalculatorServiceClient) {
	log.Print("doPrime was invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", msg.Result)
	}
}
