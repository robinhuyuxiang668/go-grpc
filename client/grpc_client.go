package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"protoc-gen/proto/hello_grpc"
)

func InitClient() hello_grpc.HelloServiceClient {
	addr := ":8080"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	//defer conn.Close()
	client := hello_grpc.NewHelloServiceClient(conn)
	return client
}

func main() {
	client := InitClient()
	result, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "枫枫",
		Message: "ok",
	})
	fmt.Println(result, err)
}
