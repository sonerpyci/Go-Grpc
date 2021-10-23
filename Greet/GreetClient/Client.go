package main

import (
	"Go/GoGrpcCourse/Greet/GreetPb"
	"context"
	"google.golang.org/grpc"
	"log"
)

type server struct {}

func main() {
	log.Println("Hello I'm a Client!")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to Connect : %v", err)
	}
	defer conn.Close()

	client := GreetPb.NewGreetServiceClient(conn)

	//fmt.Printf("Created client : %f\n", client)

	GreetUnary(client)

}

func GreetUnary(client GreetPb.GreetServiceClient) {
	request := &GreetPb.GreetRequest{
		Greeting: &GreetPb.Greeting{
			FirstName: "Soner",
			LastName: "Payci",
		},
	}
	response, err := client.Greet(context.Background(), request)
	if err != nil {
		log.Fatalf("Error while calling Greet RPC : %v", err)
	}
	log.Printf("Response returned as : %v", response.Result)
}
