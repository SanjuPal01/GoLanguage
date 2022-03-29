package main

import (
	"Basics/greet/greetpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
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
