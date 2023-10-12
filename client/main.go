package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"

	pb "github.com/AnuragProg/go-grpc-prac_4/pb"
	"google.golang.org/grpc"
)


func main(){
	reader := bufio.NewReader(os.Stdin)
	msg := ""

	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf(err.Error())
	}
	client := pb.NewChatServiceClient(conn)

	stream, _ := client.Converstion(context.TODO())
	for {
		log.Println("Enter message:")	
		msg, _ = reader.ReadString('\n')
		msg = strings.TrimFunc(msg, func(r rune)bool{
			return r == '\n' || r == '\r'
		})
		if err = stream.Send(&pb.Message{Msg: msg}); err != nil{
			log.Println(err.Error())
		}
		resp, err := stream.Recv()
		if err != nil{
			log.Println(err.Error())
		}else{
			log.Println("Received:", resp.GetMsg())
		}
	}
}
