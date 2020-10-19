package main

import (
	"context"
	"log"
	"time"

	pb "github.com/panyuenlau/mygrpc-server/proto"
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

	sendMsg := "Hello Server!"
	r, err := c.SayHello(ctx, &pb.Request{ReqeustMessage: sendMsg})

	if err != nil {
		log.Fatalf("ERROR: failed to send message to the server: %v", err)
	} else {
		log.Printf("The message \"" + sendMsg + "\"" + " was sent to the server, waiting for response...")
	}

	log.Printf("Message from the server: \"%s\n", r.GetReplyMessage()+"\"")
}
