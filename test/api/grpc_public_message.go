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
		go client.publicMessage()
	}

	client.Wg.Wait()
}

func (client *GRPCClient) publicMessage() {
	defer client.Wg.Done()

	c := pb.NewKafkaClient(client.Conn)

	for i := 0; i < 100000; i++ {
		_, err := c.PublicMessage(context.Background(), &pb.PublicMessageRequest{
			TopicName: "big_john",
			Data:      []byte("hello world" + strconv.Itoa(i)),
		})

		if err != nil {
			logrus.Errorf("kafka client public message error: %v", err)
			continue
		}
	}

}
