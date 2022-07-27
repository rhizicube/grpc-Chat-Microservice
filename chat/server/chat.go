package main

import (
	"bufio"
	pb "gRPC/chat/proto"
	"log"
	"os"
)

func (s *Server) Chat(stream pb.ChatServive_ChatServer) error {
	for {
		req, err := stream.Recv()

		if err != nil {
			log.Fatalf("Error while receiving client stream: %v", err)
		}
		log.Printf("BOB: %s", req.MessageReq)

		go func() {
			for {
				sc := bufio.NewReader(os.Stdin)
				input, _ := sc.ReadString('\n')

				err := stream.Send(&pb.ChatResponse{
					MessageRes: input,
				})
				if err != nil {
					log.Fatalf("Error while sending data to client: %v\n", err)
				}
			}
		}()
	}
}
