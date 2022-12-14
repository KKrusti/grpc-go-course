package main

import (
	"context"
	pb "github.com/Clement-Jean/grpc-go-course/calculator/proto"
	"log"
)

func (s *Server) Sum(ctx context.Context, in *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	log.Printf("Sum was invoked with %v\n", in)
	return &pb.CalculateResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
