package main

import (
	pb "GoProgrammingJourney/gorpc/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, _ := grpc.Dial(":8000", grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	err := SayHello(client)
	if err != nil {
		log.Fatalf("SayHello err: %v", err)
	}

	err = SayList(client, &pb.HelloRequest{
		Name: "SayList",
	})
	if err != nil {
		log.Fatalf("SayList err: %v", err)
	}

	err = SayRecord(client, &pb.HelloRequest{
		Name: "SayRecord",
	})
	if err != nil {
		log.Fatalf("SayRecord err: %v", err)
	}

	err = SayRoute(client, &pb.HelloRequest{
		Name: "SayRoute",
	})
	if err != nil {
		log.Fatalf("SayRoute err: %v", err)
	}
}
