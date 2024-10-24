package grpc

import (
	"context"

	"github.com/guimochila/microservices-with-go/internal/grpcutil"
	moviev1 "github.com/guimochila/microservices-with-go/pkg/api/movie"
	"github.com/guimochila/microservices-with-go/pkg/discovery"
	"github.com/guimochila/microservices-with-go/rating/pkg/model"
)

// Getway defines a gRPC gateway for rating service.
type Gateway struct {
	registry discovery.Registry
}

// New creates a new gRPC gateway for a rating service.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

// GetAggregatedRating returns the aggregated rating for a record or ErrNotfound
// if there are not rating for it.
func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID, recordType model.RecordType) (float64, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "rating", g.registry)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	client := moviev1.NewRatingServiceClient(conn)
	resp, err := client.GetAggregatedRating(ctx, &moviev1.GetAggregatedRatingRequest{
		RecordId:   string(recordID),
		RecordType: string(recordType),
	})
	if err != nil {
		return 0, err
	}
	return resp.RatingValue, nil
}
