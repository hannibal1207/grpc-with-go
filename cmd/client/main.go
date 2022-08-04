package main

import (
	"context"
	"log"

	"exemplo.com/api/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewSendMessageClient(conn)

	req := &pb.Request{
		Message: "Hello Grpc",
	}

	res, err := client.RequestMessage(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res.GetStatus())
}
