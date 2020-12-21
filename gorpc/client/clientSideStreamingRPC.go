package main

import (
	pb "GoProgrammingJourney/gorpc/proto"
	"context"
	"log"
)

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayRecord(context.Background())
	if err != nil {
		return err
	}
	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Println("resp.SayRecord: ", resp)
	return nil
}
