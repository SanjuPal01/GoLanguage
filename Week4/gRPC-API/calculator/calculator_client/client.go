package main

import (
	"Basics/calculator/calculatorpb"
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
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// DoUnary(c)

	// DoServerStreaming(c)

	// DoClientStreaming(c)

	// DoBiDiStreaming(c)

	DoErrorUnaryCall(c)
}

func DoUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Unary RPC")
	req := &calculatorpb.SumRequest{
		Data: &calculatorpb.Data{
			FirstNum:  3,
			SecondNum: 10,
		},
	}
	resp, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling the Sum Function: %v", err)
	}
	fmt.Printf("Response from sum: %v\n", resp.Result)
}

func DoServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Server Streaming RPC")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		// Num: 12390392840,
		Num: 120,
	}
	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling the PrimeNumberDecompostion Method: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// Reached the end of stream
			break
		} else if err != nil {
			log.Fatalf("Error while fetching the data from stream: %v", err)
		}
		fmt.Printf("Response from PrimeNumberDecomposition: %v\n", msg.GetRes())
	}
}

func DoClientStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Client Streamig RPC")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while Opening Stream: %v", err)
	}
	numbers := []int32{3, 5, 9, 54, 23}
	for _, number := range numbers {
		fmt.Println("Sending Number: ", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Num: number,
		})
		time.Sleep(1 * time.Second) // for the streaming effect
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while Receiving Response: %v", err)
	}
	fmt.Println("The Average is ", resp.GetRes())
}

func DoBiDiStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a BiDi Streamig RPC")
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while calling the FindMaximum: %v", err)
	}
	numbers := []int64{1, 5, 3, 6, 2, 20}
	ch := make(chan struct{})
	go func() {
		for _, num := range numbers {
			fmt.Printf("Sending number: %v\n", num)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Num: num,
			})
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
				log.Fatalf("Problem while reading server stream: %v", err)
				break
			}
			fmt.Println("Received a new maximum:", resp.GetCurr_Max())
		}
		close(ch)
	}()
	<-ch
}

func DoErrorUnaryCall(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Error Unary RPC")
	// correct call
	DoErrorCall(c, 10)

	// error call
	DoErrorCall(c, -2)

}

func DoErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	resp, err := c.SquareRoot(context.Background(), &calculatorpb.SquareRootRequest{
		Number: n,
	})
	if err != nil {
		respError, ok := status.FromError(err)
		if ok {
			// actual error from gRPC(user error)
			fmt.Printf("Error Message from server: %v", respError.Message())
			fmt.Println(respError.Code())
			if respError.Code() == codes.InvalidArgument {
				fmt.Println("You probably sent a negative number!")
				return
			}
		} else {
			// error not raised by gRPC
			log.Fatalf("Big Error calling SquareRoot: %v", err)
			return
		}
	}
	fmt.Printf("Result of Square root of %v: %v\n\n", n, resp.GetNumberRoot())
}
