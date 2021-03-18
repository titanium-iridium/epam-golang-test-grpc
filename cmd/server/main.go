package main

import (
	"context"
	"fmt"
	"github.com/titanium-iridium/epam-golang-test-grpc/pkg/common"
	pb "github.com/titanium-iridium/epam-golang-test-grpc/pkg/test"
	"google.golang.org/grpc"
	"net"
	"os"
)

type server struct {
	pb.UnimplementedProducerServer
	channel chan *common.Message
}

func main() {
	config := common.GetConfig()
	listener, err := net.Listen("tcp", config.Address)
	if err != nil {
		common.LogError("Network failure", err)
		os.Exit(1)
	}

	fmt.Println("Server listener started at " + config.Address)

	msgChannel := make(chan *common.Message)
	go common.ConsoleOutput(msgChannel)

	grpcServer := grpc.NewServer()
	server := server{channel: msgChannel}
	pb.RegisterProducerServer(grpcServer, &server)

	err = grpcServer.Serve(listener)
	if err != nil {
		common.LogError("Server failure", err)
		os.Exit(1)
	}
}

// Server method implementation
func (srv *server) SaySome(_ context.Context, in *pb.Request) (*pb.Response, error) {
	srv.channel <- &common.Message{Time: in.Time.AsTime(), Text: in.Text}
	return &pb.Response{Ok: true}, nil
}
