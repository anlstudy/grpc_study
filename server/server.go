package main

import (
	"google.golang.org/grpc"
	pb "grpc_demo/proto"
	"grpc_demo/server/server"
	"log"
	"net"
)

const (
	Network = "tcp"
	Address = ":8000"
)

func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &server.SimpleService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
