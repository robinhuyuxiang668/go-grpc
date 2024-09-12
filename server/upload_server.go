package main

import (
	"bufio"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"protoc-gen/proto/upload_grpc"
)

type ClientStream struct{}

func (ClientStream) UploadFile(stream upload_grpc.UploadService_UploadFileServer) error {

	file, err := os.OpenFile("static/copy-golden-dark.mp4", os.O_CREATE|os.O_WRONLY, 0600)
	fmt.Println("file ", file, err)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	var index int
	for {
		response, err := stream.Recv()
		fmt.Println("serve read err== io.EOF ", err == io.EOF)

		if err == io.EOF {
			break
		}

		index++
		writer.Write(response.Content)
		fmt.Printf("serve read 第%d 次， 写入 %d 数据\n", index, len(response.Content))
	}
	writer.Flush()
	stream.SendAndClose(&upload_grpc.Response{Text: "完毕了"})
	fmt.Println("server UploadFile done", file)

	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	upload_grpc.RegisterUploadServiceServer(server, &ClientStream{})

	server.Serve(listen)
}
