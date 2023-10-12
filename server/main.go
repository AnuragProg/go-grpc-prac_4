package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/AnuragProg/go-grpc-prac_4/pb"
	"google.golang.org/grpc"
)


const PORT = 3000

type Server struct {
	pb.UnimplementedChatServiceServer
}

func (s *Server) Converstion(cs pb.ChatService_ConverstionServer) error {
	responses := map[string]string{
		"hello":    "greetings",
		"goodbye":  "farewell",
		"happy":    "joyful",
		"sad":      "melancholy",
		"go":       "golang",
		"pizza":    "delicious",
		"sun":      "sunny",
		"rain":     "wet",
		"coffee":   "awake",
		"book":     "reading",
	}
	for {
		msg, err := cs.Recv()
		if err != nil{
			log.Println("got error:", err.Error())			
			return err
		}
		log.Println("got message from client:", msg.GetMsg())
		log.Println("got message from client:", msg.GetMsg())

		if resp, ok := responses[msg.GetMsg()]; ok {
			cs.Send(&pb.Message{
				Msg: resp,
			})
		}else{
			cs.Send(&pb.Message{
				Msg: "I have no proper response",
			})
		}
	}
}

func main(){
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v",PORT))
	if err != nil{
		log.Fatalf("unable to listen on port %v because of %v", PORT, err.Error())
	}
	
	server := grpc.NewServer()
	pb.RegisterChatServiceServer(server, &Server{})
	log.Println("Listening on", PORT)
	if err = server.Serve(lis); err != nil{
		log.Fatalf("unable to serve grpc server")
	}
}
