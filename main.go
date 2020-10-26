package main

import (
	"context"
	"log"
	"os"
	"time"

	"net/http"

	pb "github.com/panyuenlau/mygrpc-client/proto"
	"google.golang.org/grpc"
)

const (
	connectionDeadline = 5 // max amount of time(in seconds) the client tries to build connection with the server
)

/*
1. If the client cannot connect, logs error, return
2. If the client connects, but the server doesn't respond before timeout, logs error, return
*/

func handler(w http.ResponseWriter, req *http.Request) {
	log.Println("Check server status request received")

	address := os.Getenv("SERVICE_TO_GRPC_SERVER")

	if len(address) == 0 {
		log.Println("ERROR: failed to retreive the service name to reach the gRPC server, existing...")
		return
	}

	log.Println("Now dialing the gRPC server through", address)

	// set timeout for dialing to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Duration(connectionDeadline)*time.Second))

	if err != nil {
		log.Println("ERROR: failed to connect gRPC server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("gRPC server connected!")

	defer conn.Close() // defer the execution of a function until the surrounding functions return

	c := pb.NewGreetingClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	sendMsg := "Hello Server! " + time.Now().Format("2006-01-02 15:04:05 PM")
	r, err := c.SayHello(ctx, &pb.Request{ReqeustMessage: sendMsg})

	if err != nil {
		log.Println("ERROR: failed to send message to the server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("The message \"%s\" has been sent to the server", sendMsg)
	log.Printf("Respopnse from the server: \"%s\"", r.GetReplyMessage())
}

func main() {
	http.HandleFunc("/serverstatus", handler)
	err := http.ListenAndServe(":8081", nil)
	log.Fatal("ListenAndServe: ", err)
}
