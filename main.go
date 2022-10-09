package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/wan-maoyuan/kafka-go/pkg/server"
)

var (
	Name       string = "kafka-go server"
	Version    string //版本
	CommitHash string //git 提交的 hash 值
	BuildTime  string //编译时间
)

func init() {
	showVersion()
}

func main() {
	svr, err := server.NewKafkaServer()
	if err != nil {
		logrus.Fatalf("create kafka server error: %v", err)
	}

	go func() {
		if err := svr.Run(); err != nil {
			logrus.Fatalf("run kafka server error: %v", err)
		}
	}()

	stopSingal := make(chan os.Signal, 1)
	signal.Notify(stopSingal, syscall.SIGINT, syscall.SIGTERM)
	<-stopSingal

	svr.Stop()
}

func showVersion() {
	fmt.Printf(`
Name: %s
Version: %s
CommitHash: %s
BuildTime: %s

`, Name, Version, CommitHash, BuildTime)
}
