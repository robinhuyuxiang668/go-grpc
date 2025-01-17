package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"os"
	"protoc-gen/proto/upload_grpc"
)

func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
	}
	defer conn.Close()
	// 初始化客户端
	client := upload_grpc.NewUploadServiceClient(conn)
	stream, err := client.UploadFile(context.Background())

	file, err := os.Open("static/golden-dark.mp4")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for {
		buf := make([]byte, 2048)
		_, err = file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		stream.Send(&upload_grpc.FileRequest{
			FileName: "golden-dark.mp4",
			Content:  buf,
		})
	}
	response, err := stream.CloseAndRecv()
	fmt.Println(response, err)
}
