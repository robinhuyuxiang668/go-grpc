package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"protoc-gen/proto/hello_grpc"
)

// HelloService 得有一个结构体，需要实现这个服务的全部方法,叫什么名字不重要
type HelloService struct {
}

func (HelloService) SayHello(ctx context.Context, request *hello_grpc.HelloRequest) (response *hello_grpc.HelloResponse, err error) {
	fmt.Println("入参：", request.Name, request.Message)
	response = new(hello_grpc.HelloResponse)
	response.Name = "你好"
	response.Message = "ok"
	return
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 创建一个gRPC服务器实例。
	server := grpc.NewServer()
	service := HelloService{}
	// 将service结构体注册为gRPC服务。
	hello_grpc.RegisterHelloServiceServer(server, &service)
	fmt.Println("grpc server running :8080")
	// 开始处理客户端请求。
	err = server.Serve(listen)
}
