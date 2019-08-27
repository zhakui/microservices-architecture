package main

import (
	"context"
	"fmt"
	"os"

	"product-service/protocol/rpc"
	"product-service/service"
	"product-service/protocol/rest"
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
