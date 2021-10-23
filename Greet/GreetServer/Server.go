package main

import (
	"Go/GoGrpcCourse/Greet/GreetPb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {}
func (*server) Greet(ctx context.Context, req *GreetPb.GreetRequest) (*GreetPb.GreetResponse, error) {
	fmt.Printf("greet function was invoked with %v on the server side", req)

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName + " " + lastName
	res := &GreetPb.GreetResponse{
		Result: result,
	}
	return res, nil
}


func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	s := grpc.NewServer()
	GreetPb.RegisterGreetServiceServer(s, &server{})

	if err:= s.Serve(lis); err!= nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
