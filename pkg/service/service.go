package service

import (
	"context"
	"os"
	"path/filepath"

	pb "github.com/wan-maoyuan/kafka-go/api/kafka"
	"github.com/wan-maoyuan/kafka-go/pkg/topic"
	"github.com/wan-maoyuan/kafka-go/pkg/utils"
)

type KafkaService struct {
	pb.UnimplementedKafkaServer
	topicCtl *topic.Topic
}

func init() {
	if _, err := os.Stat(utils.C.Log.FilePath); err != nil {
		os.Mkdir(utils.C.Log.FilePath, 0755)
	}
}

func NewKafkaService() (*KafkaService, error) {
	t, err := topic.NewTopic(filepath.Join(utils.C.Log.FilePath, "topic"))
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
	topics := make([]*pb.GetAllTopicsResponse_Topic, 0)
	for _, topic := range s.topicCtl.GetAll() {
		topics = append(topics, &pb.GetAllTopicsResponse_Topic{
			Name: topic,
		})
	}

	return &pb.GetAllTopicsResponse{
		Count:  uint32(len(topics)),
		Topics: topics,
	}, nil
}

func (s *KafkaService) PublicMessage(svr pb.Kafka_PublicMessageServer) error {
	// for {
	// 	resp, err := svr.Recv()
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (s *KafkaService) SubscribeMessage(req *pb.SubscribeMessageRequest, svr pb.Kafka_SubscribeMessageServer) error {
	return nil
}
