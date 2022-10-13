package api

import (
	"context"
	"strconv"

	"github.com/sirupsen/logrus"
	pb "github.com/wan-maoyuan/kafka-go/api/kafka"
)

func (client *GRPCClient) TestPublicMessage() {
	for i := 0; i < client.Thread; i++ {
		client.Wg.Add(1)
		go client.publicMessageThread()
	}

	client.Wg.Wait()
}

func (client *GRPCClient) publicMessageThread() {
	defer client.Wg.Done()

	for i := 0; i < client.Count; i++ {
		client.publicMessage()
	}
}

func (client *GRPCClient) publicMessage() {
	c := pb.NewKafkaClient(client.Conn)

	for i := 0; i < 1000; i++ {
		_, err := c.PublicMessage(context.Background(), &pb.PublicMessageRequest{
			TopicName: "big_john",
			Data:      []byte("hello world" + strconv.Itoa(i)),
		})
		if err != nil {
			logrus.Errorf("kafka client public message error: %v", err)
			return
		}
	}

}
