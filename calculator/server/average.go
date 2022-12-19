package main

import (
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"io"
	"log"
)

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average was invoked")
	var sum float32 = 0
	var counter float32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{
				Result: sum / counter,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving req: %v\n", req)
		sum += float32(req.Number)
		counter++
	}

	return nil
}
