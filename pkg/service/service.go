package service

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	pb "github.com/wan-maoyuan/kafka-go/api/kafka"
	"github.com/wan-maoyuan/kafka-go/pkg/storage"
	"github.com/wan-maoyuan/kafka-go/pkg/utils"
)

type KafkaService struct {
	pb.UnimplementedKafkaServer
	storageManage *storage.Storage
}

// 初始化数据文件夹，不存在就新建
func initDataDir() {
	if _, err := os.Stat(utils.C.Log.FilePath); err != nil {
		logrus.Errorf("data dir stat is error: %v", err)

		if err := os.Mkdir(utils.C.Log.FilePath, 0755); err != nil {
			logrus.Errorf("mkdir data dir error: %v", err)
		}
	}
}

func NewKafkaService() (*KafkaService, error) {
	initDataDir()

	s, err := storage.NewSorage()
	if err != nil {
		return nil, err
	}

	return &KafkaService{
		storageManage: s,
	}, nil
}

func (s *KafkaService) CreateTopic(_ context.Context, req *pb.CreateTopicRequest) (*pb.CreateTopicResponse, error) {

	return &pb.CreateTopicResponse{}, nil
}

func (s *KafkaService) DeleteTopic(_ context.Context, req *pb.DeleteTopicRequest) (*pb.DeleteTopicResponse, error) {

	return &pb.DeleteTopicResponse{}, nil
}

func (s *KafkaService) GetAllTopics(_ context.Context, req *pb.GetAllTopicsRequest) (*pb.GetAllTopicsResponse, error) {

	return &pb.GetAllTopicsResponse{}, nil
}

func (s *KafkaService) PublicMessage(_ context.Context, req *pb.PublicMessageRequest) (*pb.PublicMessageResponse, error) {
	if err := s.storageManage.SaveMessage(req.TopicName, req.Data); err != nil {
		return &pb.PublicMessageResponse{
			IsPublic: false,
			Message:  err.Error(),
		}, err
	}

	return &pb.PublicMessageResponse{
		IsPublic: true,
		Message:  "",
	}, nil
}

func (s *KafkaService) SubscribeMessage(req *pb.SubscribeMessageRequest, svr pb.Kafka_SubscribeMessageServer) error {
	offset := req.Offset
	for {
		logrus.Debugf("SubscribeMessage api offset: %d, topic: %s", offset, req.TopicName)
		msg, err := s.storageManage.GetMessage(req.TopicName, offset)
		if err != nil {
			logrus.Errorf("SubscribeMessage api get message error: %v", err)
			svr.Context().Done()
			return err
		}

		if err := svr.Send(&pb.SubscribeMessageResponse{
			Data:   msg,
			Offset: offset,
		}); err != nil {
			logrus.Errorf("SubscribeMessage api send message error: %v", err)
			svr.Context().Done()
			return err
		}

		offset += 1
	}
}
