package server

import (
	"fmt"
	"net"

	pb "github.com/wan-maoyuan/kafka-go/api/kafka"
	"github.com/wan-maoyuan/kafka-go/pkg/service"
	"google.golang.org/grpc"
)

type KafkaServer struct {
	lis net.Listener
	gs  *grpc.Server
}

func NewKafkaServer() (*KafkaServer, error) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return nil, fmt.Errorf("creater tcp listener error: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterKafkaServer(grpcServer, &service.KafkaService{})

	return &KafkaServer{
		lis: listener,
		gs:  grpcServer,
	}, nil
}

func (server *KafkaServer) Run() error {
	if err := server.gs.Serve(server.lis); err != nil {
		return fmt.Errorf("run grpc server error: %v", err)
	}

	return nil
}

func (server *KafkaServer) Stop() error {
	server.gs.Stop()

	return server.lis.Close()
}
