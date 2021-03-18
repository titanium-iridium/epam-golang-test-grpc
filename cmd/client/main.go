package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/titanium-iridium/epam-golang-test-grpc/pkg/common"
	pb "github.com/titanium-iridium/epam-golang-test-grpc/pkg/test"
	"google.golang.org/grpc"
	"os"
	"time"
)

func main() {
	// customize logger
	//log.SetFlags(0)
	//log.SetOutput(new(common.LogWriter))

	config := common.GetConfig()

	msgChannel := make(chan string)
	fmt.Println("Client started")
	go common.ConsoleInput(msgChannel)
	for text := range msgChannel {
		send(text, config.Address)
	}
}

func send(text, address string) {
	dialContext, dialCancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer dialCancel()

	client, err := grpc.DialContext(
		dialContext,
		address,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		common.LogError("Dialing failure", err)
		os.Exit(1)
	}
	defer func() {
		_ = client.Close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c := pb.NewProducerClient(client)
	_, err = c.SaySome(ctx, &pb.Request{Text: text, Time: ptypes.TimestampNow()})
	if err != nil {
		common.LogError("Sending failure", err)
	}
}
