package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"protoc-gen/proto/download_grpc"
)

type ServiceStream struct{}

func (ServiceStream) DownLoadFile(request *download_grpc.Request, stream download_grpc.DownLoadService_DownLoadFileServer) error {
	fmt.Println(request)
	file, err := os.Open("static/golden-dark.mp4")
	if err != nil {
		return err
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
		stream.Send(&download_grpc.FileResponse{
			Content: buf,
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
	download_grpc.RegisterDownLoadServiceServer(server, &ServiceStream{})

	server.Serve(listen)
}
