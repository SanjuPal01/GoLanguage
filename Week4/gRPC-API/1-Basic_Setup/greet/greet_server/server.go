package main

import (
	"Basics/greet/greetpb"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

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

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with: %v\n", req)
	firstName := req.GetGreet().GetFirstName()
	for i := 0; i < 10; i++ {
		res := "Hello " + firstName + " number " + strconv.Itoa(i)
		resp := &greetpb.GreetManyTimesResponse{
			Result: res,
		}
		stream.Send(resp)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet function was invoked")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// even we can invoke this method anywhere - that means we can return to client anytime in between the stream is coming
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while Reading Client Stream: %v", err)
		}
		firstName := req.GetGreet().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("GreetEveryone function was invoked")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading the client stream: %v\n", err)
			return err
		}
		firstName := req.GetGreet().GetFirstName()
		res := "Hello " + firstName + "! "
		err = stream.Send(&greetpb.GreetEveryoneResponse{
			Result: res,
		})
		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
			return err
		}
	}
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
