package main

import (
	"Basics/calculator/calculatorpb"
	"context"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Greet function was invoked with: %v\n", req)
	firstNum := req.GetData().FirstNum
	secondNum := req.GetData().SecondNum
	res := firstNum + secondNum
	resp := &calculatorpb.SumResponse{
		Result: res,
	}
	return resp, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition function was invoked with: %v\n", req)
	// var k int32 = 2
	k := int64(2)
	N := req.GetNum()
	for N > 1 {
		if N%k == 0 { // if k evenly divides into N
			resp := &calculatorpb.PrimeNumberDecompositionResponse{
				Res: k,
			}
			stream.Send(resp)           // this is a factor
			time.Sleep(1 * time.Second) // Adding sleep function for seeing the streaming effect
			N = N / k                   // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Println("ComputeAverage Function is Invoked")
	sum := int32(0)
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			avg := float64(sum) / float64(count)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Res: avg,
			})
		}
		if err != nil {
			log.Fatalf("Error while Receiving the Data from Client Stream: %v", err)
		}
		count++
		sum += req.GetNum()
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Println("ComputeAverage Function is Invoked")
	currMax := int64(math.MinInt64)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while Receiving the Stream Data: %v", err)
			return err
		}
		num := req.GetNum()
		if num > currMax {
			currMax = num
			sendErr := stream.Send(&calculatorpb.FindMaximumResponse{
				Curr_Max: currMax,
			})
			if sendErr != nil {
				log.Fatalf("Error while sending data to client: %v", sendErr)
				return sendErr
			}
		}
	}
}

func (*server) SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error) {
	fmt.Println("SquareRoot Function is Invoked")
	number := req.GetNumber()
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v\n", number),
		)
	}
	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil
}

func main() {
	fmt.Println("Starting Server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to Serve: %v", err)
	}
}
