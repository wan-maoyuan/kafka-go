package api

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	pb "github.com/wan-maoyuan/kafka-go/api/kafka"
)

func (client *GRPCClient) TestSubscribeMessage() {
	for i := 0; i < client.Thread; i++ {
		client.Wg.Add(1)
		go client.subscribeMessage()
	}

	client.Wg.Wait()
}

func (client *GRPCClient) subscribeMessage() {
	defer client.Wg.Done()

	c := pb.NewKafkaClient(client.Conn)

	stream, err := c.SubscribeMessage(context.Background(), &pb.SubscribeMessageRequest{
		TopicName: "big_john",
		Offset:    0,
	})

	if err != nil {
		logrus.Errorf("SubscribeMessage get stream client error: %v", err)
		return
	}
	defer stream.CloseSend()

	for {
		resp, err := stream.Recv()
		if err != nil {
			logrus.Errorf("stream recive message error: %v", err)
			return
		}

		fmt.Println("test client recive a message = ", string(resp.Data))
	}

}
