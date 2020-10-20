package main

import (
	"context"
	"log"
	"time"

	pb "github.com/panyuenlau/mygrpc-client/proto"
	"google.golang.org/grpc"
)

const (
	// address for local debug
	address = "localhost:50051"
	// address used in the K8s cluster
	// address = "grpc-service:50051"

	connectionDeadline = 5 // max amount of time the client tries to build connection with the server
)

/*
1. If the client cannot connect, logs error, return
2. If the client connects, but the server doesn't respond before timeout, logs error, return
*/

func main() {
	log.Println("Connecting to server...")

	// set timeout for dialing to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Duration(connectionDeadline)*time.Second))

	if err != nil {
		log.Fatalf("ERROR: failed to connect: %v", err)
		return
	}

	log.Println("Server connected!")

	defer conn.Close() // defer the execution of a function until the surrounding functions return

	for {
		c := pb.NewGreetingClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		sendMsg := "Hello Server! " + time.Now().Format("2006-01-02 15:04:05 PM")
		r, err := c.SayHello(ctx, &pb.Request{ReqeustMessage: sendMsg})

		if err != nil {
			log.Fatalf("ERROR: failed to send message to the server: %v", err)
			return
		}
		log.Printf("The message \"%s\" has been sent to the server", sendMsg)

		log.Printf("Respopnse from the server: \"%s\"\n\n", r.GetReplyMessage())
		time.Sleep(time.Second * 5)
	}
}
