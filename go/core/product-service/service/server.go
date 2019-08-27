package v1

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"product-service/api"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// The ArtManagerServer is a article manager system
type ArtManagerServer struct {
	id          int64
	title       string
	author      int64
	category    int64
	lastUpdated *timestamp.Timestamp
}

// NewArtManagerServer we can use it to create a ArtManagerServer instance.
func NewArtManagerServer() *ArtManagerServer {
	t := time.Now().In(time.UTC)
	reminder, _ := ptypes.TimestampProto(t)
	return &ArtManagerServer{1, "go-rpc-rest", 1, 1, reminder}
}

// GetArticleInfor Get the article information.
func (ams *ArtManagerServer) GetArticleInfor(ctx context.Context, req *v1.ArtRequest) (*v1.ArtRespones, error) {
	return &v1.ArtRespones{
		Api: req.Api,
		Article: &v1.Article{
			Id:          req.Id,
			Title:       ams.title,
			Author:      ams.author,
			Category:    ams.category,
			LastUpdated: ams.lastUpdated,
		},
	}, nil
}
