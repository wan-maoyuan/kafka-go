package service

import (
	"context"

	pb "github.com/wan-maoyuan/kafka-go/api/kafka"
)

type KafkaService struct {
	pb.UnimplementedKafkaServer
}

func (s *KafkaService) CreateTopic(_ context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {

	return nil, nil
}

func (s *KafkaService) DeleteTopic(_ context.Context, req *pb.DeleteTopicRequest) (*pb.DeleteTopicResponse, error) {
	return nil, nil
}

func (s *KafkaService) GetAllTopics(_ context.Context, req *pb.GetAllTopicsRequest) (*pb.GetAllTopicsResponse, error) {
	return nil, nil
}

func (s *KafkaService) PublicMessage(svr pb.Kafka_PublicMessageServer) error {
	return nil
}

func (s *KafkaService) SubscribeMessage(req *pb.SubscribeMessageRequest, svr pb.Kafka_SubscribeMessageServer) error {
	return nil
}
