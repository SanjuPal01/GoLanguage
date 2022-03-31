package main

import (
	"Basics/greet/greetpb"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Creating Client")
	// cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // because by default it want ssl certificates
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) // New Way - because old one is deprecated
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Client Created: %f\n", c)
	// DoUnary(c)

	// DoServerStreaming(c)

	// DoClientStreaming(c)

	// DoBiDiStreaming(c)

	doUnaryWithDeadline(c, 5*time.Second) // should complete
	doUnaryWithDeadline(c, 2*time.Second) // should timeout

}

func DoUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Staring to do a Unary RPC")
	req := &greetpb.GreetRequest{
		Greet: &greetpb.Greeting{
			FirstName: "Sanju",
			LastName:  "Pal",
		},
	}
	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling the Greet Function: %v", err)
	}
	fmt.Printf("Response from greet: %v\n", resp.Result)
}

func DoServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Staring to do a Server Streaming RPC")
	req := &greetpb.GreetManyTimesRequest{
		Greet: &greetpb.Greeting{
			FirstName: "Sanju",
			LastName:  "Pal",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greetManyTimes: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// it will give EOF when we reached the end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v", err)
		}
		fmt.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}
}

func DoClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Staring to do a Server Streaming RPC")
	requests := []*greetpb.LongGreetRequest{
		{
			Greet: &greetpb.Greeting{
				FirstName: "Sanju",
			},
		},
		{
			Greet: &greetpb.Greeting{
				FirstName: "Sameeer",
			},
		},
		{
			Greet: &greetpb.Greeting{
				FirstName: "Karan",
			},
		},
		{
			Greet: &greetpb.Greeting{
				FirstName: "Sahil",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while Calling longGreet: %v\n", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while Receiving the response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", resp.GetResult())
}

func DoBiDiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling the LongGreet method: %v", err)
	}

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greet: &greetpb.Greeting{
				FirstName: "Sanju",
			},
		},
		{
			Greet: &greetpb.Greeting{
				FirstName: "Sameeer",
			},
		},
		{
			Greet: &greetpb.Greeting{
				FirstName: "Karan",
			},
		},
		{
			Greet: &greetpb.Greeting{
				FirstName: "Sahil",
			},
		},
	}

	// program will not terminate until we close the channel
	// that's why we are using it hence our stream can works fine
	ch := make(chan struct{})
	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			fmt.Printf("Received: %v\n", resp.GetResult())
		}
		close(ch)
	}()

	// block until everything is done
	<-ch
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Staring to do a UnaryWithDeadline RPC")
	req := &greetpb.GreetWithDeadlineRequest{
		Greet: &greetpb.Greeting{
			FirstName: "Sanju",
			LastName:  "Pal",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		errStatus, ok := status.FromError(err)
		if ok {
			// error from grpc
			if errStatus.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout has hit! Deadline was exceeded")
			} else {
				fmt.Printf("Unexpected Error: %v", err)
			}
		} else {
			log.Fatalf("Error while calling the Greet Function: %v", err)
		}
		return
	}
	fmt.Printf("Response from greet: %v\n", resp.Result)
}
