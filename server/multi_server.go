package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"protoc-gen/multi_reference_proto/proto"
)

type VideoServer struct {
}

func (VideoServer) Look(ctx context.Context, request *proto.Request) (res *proto.Response, err error) {
	fmt.Println("video:", request)
	return &proto.Response{
		Name: "ffa Look",
	}, nil
}

type OrderServer struct {
}

func (OrderServer) Buy(ctx context.Context, request *proto.Request) (res *proto.Response, err error) {
	fmt.Println("order:", request)
	return &proto.Response{
		Name: "fengfeng Buy",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterVideoServiceServer(s, &VideoServer{})
	proto.RegisterOrderServiceServer(s, &OrderServer{})
	fmt.Println("grpc server程序运行在：8080")
	err = s.Serve(listen)
}
