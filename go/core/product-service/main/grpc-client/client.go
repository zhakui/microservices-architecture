package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	v1 "product-service/api"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(":2536", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewArticleManagementServiceClient(conn)

	ctx := context.Background()
	artRequest := &v1.ArtRequest{
		Api: "1",
		Id:  1,
	}
	artInfor, _ := c.GetArticleInfor(ctx, artRequest)
	fmt.Printf("%+v\n", artInfor)
}
