package main

import (
	"context"
	"fmt"
	"os"

	"../../../pkg/protocol"
	"../../../pkg/protocol/rest"
	"../../../pkg/service/v1"
)

func main() {
	ctx := context.Background()

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, "2536", "8080")
	}()

	server := v1.NewArtManagerServer()
	if err := grpc.RunServer(ctx, server, "2536"); err != nil {
		fmt.Fprintln(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
