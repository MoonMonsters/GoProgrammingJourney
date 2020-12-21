package main

import (
	pb "GoProgrammingJourney/gorpc/proto"
	"context"
	"log"
)

func SayHello(client pb.GreeterClient) error {
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "chentao",
	})
	if err != nil {
		return err
	}
	log.Printf("client.SayHello resp: %s\n", resp.Message)
	return nil
}
