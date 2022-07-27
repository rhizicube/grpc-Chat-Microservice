package main

import (
	"bufio"
	"context"
	pb "gRPC/chat/proto"
	"io"
	"log"
	"os"
)

func doChat(c pb.ChatServiveClient) {
	for {
		stream, err := c.Chat(context.Background())
		if err != nil {
			log.Fatalf("Error while creating stream: %v", err)
		}
		req := &pb.ChatRequest{}

		sc := bufio.NewReader(os.Stdin)
		input, _ := sc.ReadString('\n')
		req.MessageReq = input

		stream.Send(req)

		go func() {
			for {
				res, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("Error while receiving: %v\n", err)
					break
				}
				log.Printf("ALLEN: %v", res.MessageRes)
			}
		}()
	}
}
