package model

import "github.com/guimochila/microservices-with-go/metadata/pkg/model"

// MovieDetails includes movie metadata and aggregated rating.
type MovieDetails struct {
	Rating   *float64       `json:"rating,omitempty"`
	Metadata model.Metadata `json:"metadata"`
}
