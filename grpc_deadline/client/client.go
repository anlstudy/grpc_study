package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc_demo/grpc_deadline/proto"
	"log"
	"time"
)

const Address string = ":8000"

var grpcClient pb.SimpleClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()
	ctx := context.Background()
	// 建立gRPC连接
	grpcClient = pb.NewSimpleClient(conn)
	route(ctx, 2) //2s超时
}

func route(ctx context.Context,deadlines time.Duration){
	clientDeadline := time.Now().Add(time.Duration(deadlines * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
	req:= pb.SimpleRequest{
		Data: "grpc",
	}
	res,err:=grpcClient.Route(ctx,&req)
	if err != nil {
		//获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				log.Fatalln("Route timeout!")
			}
		}
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res.Value)
}
