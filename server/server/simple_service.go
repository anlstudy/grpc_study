package main

import (
	"context"
	"fmt"
	pb "grpc_demo/proto"
	"strconv"
	"time"
)

type SimpleService struct{}


func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	fmt.Println(ctx)
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

func (s *SimpleService) ListValue(req *pb.StreamRequest, srv pb.Simple_ListValueServer) error {
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

