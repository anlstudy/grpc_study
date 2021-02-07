package server

import (
	"context"
	pb "grpc_demo/proto"
	"io"
	"log"
	"strconv"
	"time"
)

type SimpleService struct{}


func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
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

func (s *SimpleService) RouteList(srv pb.Simple_RouteListServer) error {
	for {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			//发送结果，并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res.Data)
	}
}


func (s *SimpleService) Conversations(srv pb.Simple_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&pb.StreamResponse{
			StreamValue: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Data,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Data)
	}
}
