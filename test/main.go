package main

import (
	"github.com/wan-maoyuan/kafka-go/test/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	gc := api.GRPCClient{
		Addr:   "127.0.0.1:8080",
		Thread: 1,
		Count:  1,
	}

	conn, err := grpc.Dial(gc.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("create grpc conn error: " + err.Error())
	}
	gc.Conn = conn

	gc.TestPublicMessage()
}
