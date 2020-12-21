package main

import (
	pb "GoProgrammingJourney/gorpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type GreeterServer struct {
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":8000")
	server.Serve(lis)
}

// Unary RBP: 一元RPC
func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("receive.SayHello: %s", r.Name)
	return &pb.HelloReply{Message: "hello.world"}, nil
}

// Server-side streamRPC: 服务端流式RPC
func (s *GreeterServer) SayList(r *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	fmt.Println("receive.SayList: ", r.Name)
	for n := 0; n <= 6; n++ {
		err := stream.Send(&pb.HelloReply{
			Message: "hello.list",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{
				Message: "hello.record",
			})
		}
		if err != nil {
			return err
		}

		log.Println("resp.SayRecord: ", resp)
	}
}

func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		err := stream.Send(&pb.HelloReply{
			Message: "say.route",
		})
		if err != nil {
			return err
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		n++
		log.Println("resp: ", resp)
	}
}
