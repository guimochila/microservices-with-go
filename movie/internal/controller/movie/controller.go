package movie

import (
	"context"
	"errors"

	metadataModel "github.com/guimochila/microservices-with-go/metadata/pkg/model"
	"github.com/guimochila/microservices-with-go/movie/internal/gateway"
	"github.com/guimochila/microservices-with-go/movie/pkg/model"
	ratingModel "github.com/guimochila/microservices-with-go/rating/pkg/model"
)

// ErrNotFound is returned when the movie metadata is not found.
var ErrNotFound = errors.New("movie metadata not found")

type ratingGateway interface {
	GetAggregatedRating(ctx context.Context, recordID ratingModel.RecordID, recordType ratingModel.RecordType) (float64, error)
	PutRating(ctx context.Context, recordID ratingModel.RecordID, recordType ratingModel.RecordType, rating *ratingModel.Rating) error
}

type metadataGeteway interface {
	Get(ctx context.Context, id string) (*metadataModel.Metadata, error)
}

// Controller defines a movie service controller.
type Controller struct {
	ratingGateway   ratingGateway
	metadataGeteway metadataGeteway
}

// New creates a new movie service controller.
func New(ratingGateway ratingGateway, metadataGeteway metadataGeteway) *Controller {
	return &Controller{ratingGateway, metadataGeteway}
}

// Get returns the movie details including the aggregated rating and movie
// metadata.
func (c *Controller) Get(ctx context.Context, id string) (*model.MovieDetails, error) {
	metadata, err := c.metadataGeteway.Get(ctx, id)
	if err != nil && errors.Is(err, gateway.ErrNotFound) {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	details := &model.MovieDetails{Metadata: *metadata}
	rating, err := c.ratingGateway.GetAggregatedRating(ctx, ratingModel.RecordID(id), ratingModel.RecordTypeMovie)
	if err != nil && !errors.Is(err, gateway.ErrNotFound) {
		// Just proceed in this case, it's ok not to have rating
	} else if err != nil {
		return nil, err
	} else {
		details.Rating = &rating
	}

	return details, nil
}
