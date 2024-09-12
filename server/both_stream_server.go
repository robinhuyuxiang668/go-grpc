package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"protoc-gen/proto/both_stream_grpc"
)

type BothStream struct{}

func (BothStream) Chat(stream both_stream_grpc.BothStream_ChatServer) error {
	for i := 0; i < 10; i++ {
		request, _ := stream.Recv()
		fmt.Println(request)
		stream.Send(&both_stream_grpc.Response{
			Text: "你好服务端收到了",
		})
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	both_stream_grpc.RegisterBothStreamServer(server, &BothStream{})

	server.Serve(listen)
}
