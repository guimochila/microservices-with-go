package grpcutil

import (
	"context"
	"math/rand"

	"github.com/guimochila/microservices-with-go/pkg/discovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ServiceConnection attempts to select a random service instance and returs a
// gRPC connection to it.
func ServiceConnection(ctx context.Context, serviceName string, registry discovery.Registry) (*grpc.ClientConn, error) {
	addrs, err := registry.ServiceAddress(ctx, serviceName)
	if err != nil {
		return nil, err
	}
	return grpc.NewClient(addrs[rand.Intn(len(addrs))], grpc.WithTransportCredentials(insecure.NewCredentials()))
}
