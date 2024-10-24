package grpc

import (
	"context"
	"errors"

	"github.com/guimochila/microservices-with-go/metadata/internal/controller/metadata"
	"github.com/guimochila/microservices-with-go/metadata/pkg/model"
	moviev1 "github.com/guimochila/microservices-with-go/pkg/api/movie"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines a movie metadata gRPC handler.
type Handler struct {
	moviev1.UnimplementedMetadataServiceServer
	ctrl *metadata.Controller
}

// New creates a new movie metadata gRPC handler.
func New(ctrl *metadata.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetMetadata returns movie metadata by id.
func (h *Handler) GetMetadata(ctx context.Context, req *moviev1.GetMetadataRequest) (*moviev1.GetMetadataResponse, error) {
	if req == nil || req.MovieId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}

	m, err := h.ctrl.Get(ctx, req.MovieId)
	if err != nil && errors.Is(err, metadata.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &moviev1.GetMetadataResponse{
		Metadata: model.MetadataToProto(m),
	}, nil
}
