package main

import (
	"Basics/greet/greetpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

// Implementing the Interface function
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with: %v\n", req)
	firstName := req.GetGreet().GetFirstName()
	res := "Hello, " + firstName
	resp := &greetpb.GreetResponse{
		Result: res,
	}
	return resp, nil
}

func main() {
	fmt.Println("Starting Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051") // By default grpc works on this port
	if err != nil {
		log.Fatalf("Failed to Listen: %v\n", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}
