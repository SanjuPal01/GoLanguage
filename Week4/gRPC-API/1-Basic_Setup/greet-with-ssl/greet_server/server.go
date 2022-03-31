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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
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
func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline function was invoked with: %v\n", req)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			// the client canceled the request
			fmt.Println("The Client Canceled the request!")
			return nil, status.Error(codes.Canceled, "The Client Canceled the request")
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreet().GetFirstName()
	res := "Hello, " + firstName
	resp := &greetpb.GreetWithDeadlineResponse{
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

	opts := []grpc.ServerOption{}
	// First go to greet-with-ssl directory then run it
	tls := true // main option for including or excluding ssl
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)
		if sslErr != nil {
			log.Fatalf("Failed Loading Certificates: %v", sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v\n", err)
	}
}
