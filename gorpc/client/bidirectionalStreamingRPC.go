package main

import (
	pb "GoProgrammingJourney/gorpc/proto"
	"context"
	"io"
	"log"
)

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, err := client.SayRoute(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <= 6; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Println("resp.SayRoute: ", resp)
	}

	err = stream.CloseSend()
	if err != nil {
		return err
	}

	return nil
}
