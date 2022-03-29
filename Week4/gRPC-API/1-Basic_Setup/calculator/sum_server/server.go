package main

import (
	"Basics/calculator/sumpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Greet function was invoked with: %v\n", req)
	firstNum := req.GetData().FirstNum
	secondNum := req.GetData().SecondNum
	res := firstNum + secondNum
	resp := &sumpb.SumResponse{
		Result: res,
	}
	return resp, nil
}

func main() {
	fmt.Println("Starting Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()
	sumpb.RegisterSumServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to Serve: %v", err)
	}
}
