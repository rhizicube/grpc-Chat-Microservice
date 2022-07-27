package main

import (
	"log"

	pb "gRPC/chat/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	con, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	log.Println("Client is connected with server")
	defer con.Close()

	c := pb.NewChatServiveClient(con)
	doChat(c)
}
