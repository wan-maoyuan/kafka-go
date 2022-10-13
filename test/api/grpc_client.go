package api

import (
	"sync"

	"google.golang.org/grpc"
)

type GRPCClient struct {
	Addr   string
	Thread int
	Count  int
	Wg     sync.WaitGroup
	Conn   *grpc.ClientConn
}
