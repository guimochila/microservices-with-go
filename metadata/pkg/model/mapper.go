package model

import moviev1 "github.com/guimochila/microservices-with-go/pkg/api/movie"

// MetadataToProto converts a Metadata struct into a generated proto
// counterpart.
func MetadataToProto(m *Metadata) *moviev1.Metadata {
	return &moviev1.Metadata{
		Id:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
	}
}

// MetadataFromProto converts a generated proto counterpart into a Metadata
// struct.
func MetadataFromProto(m *moviev1.Metadata) *Metadata {
	return &Metadata{
		ID:          m.Id,
		Title:       m.Title,
		Description: m.Description,
		Director:    m.Director,
	}
}
