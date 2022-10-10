package service

import (
	"context"
	"os"
	"path/filepath"

	pb "github.com/wan-maoyuan/kafka-go/api/kafka"
	"github.com/wan-maoyuan/kafka-go/pkg/topic"
)

var (
	dataDir       = "./data"
	topicFilePath = filepath.Join(dataDir, "topic")
)

type KafkaService struct {
	pb.UnimplementedKafkaServer
	topicCtl *topic.Topic
}

func init() {
	if _, err := os.Stat(dataDir); err != nil {
		os.Mkdir(dataDir, 0755)
	}
}

func NewKafkaService() (*KafkaService, error) {
	t, err := topic.NewTopic(topicFilePath)
	if err != nil {
		return nil, err
	}

	return &KafkaService{
		topicCtl: t,
	}, nil
}

func (s *KafkaService) CreateTopic(_ context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {
	s.topicCtl.Create(req.TopicName)

	return &pb.CreateTopicResponse{
		IsCreated: true,
		Message:   "",
	}, nil
}

func (s *KafkaService) DeleteTopic(_ context.Context, req *pb.DeleteTopicRequest) (*pb.DeleteTopicResponse, error) {
	s.topicCtl.Delete(req.TopicName)

	return &pb.DeleteTopicResponse{
		IsDeleted: true,
		Message:   "",
	}, nil
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
