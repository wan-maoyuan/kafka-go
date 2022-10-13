package api

import (
	"sync"

	"google.golang.org/grpc"
)

type GRPCClient struct {
	Addr   string
	Thread int
	Wg     sync.WaitGroup
	Conn   *grpc.ClientConn
}
