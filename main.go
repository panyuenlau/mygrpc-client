package main

import (
	"context"
	"log"
	"time"

	pb "./proto"
	"google.golang.org/grpc"
)

const address = "localhost:50051"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("ERROR: failed to connect: %v", err)
	}
	defer conn.Close() // defer the execution of a function until the surrounding functions return
	c := pb.NewGreetingClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.Request{ReqeustMessage: "Hello Server!"})
	if err != nil {
		log.Fatalf("ERROR: failed to send hello message to the server: %v", err)
	}

	log.Printf("Greeting: %s\n", r.GetReplyMessage())
}
