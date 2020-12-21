package main

import (
	pb "GoProgrammingJourney/gorpc/proto"
	"context"
	"io"
	"log"
)

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayList(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Println("Reply.SayList: ", resp)
	}

	return nil

}
