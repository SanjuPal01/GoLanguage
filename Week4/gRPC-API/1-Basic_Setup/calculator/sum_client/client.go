package main

import (
	"Basics/calculator/sumpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Creating Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := sumpb.NewSumServiceClient(cc)

	doUnary(c)
}

func doUnary(c sumpb.SumServiceClient) {
	fmt.Println("Starting to do a Unary RPC")
	req := &sumpb.SumRequest{
		Data: &sumpb.Data{
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
