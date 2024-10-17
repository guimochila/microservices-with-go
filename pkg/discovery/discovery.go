package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Registry defines a service registry.
type Registry interface {
	// Register creates a service instance record in the registry.
	Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error
	// Deregister removes a service instance record from the registry.
	Deregister(ctx context.Context, instanceID string, serviceName string) error
	// ServiceAddress returns the list of address of active
	// instances of the given service.
	ServiceAddress(ctx context.Context, serviceName string) ([]string, error)
	// ReportHealthyState is a push mechanism for reporting healthy state
	// to the registry
	ReportHealthyState(instanceID string, serviceName string) error
}

// ErrNotFound is returned when no service address is found.
var ErrNotFound = errors.New("no service address found")

// GenerateInstanceID generates a pseudo-random  service instance identifier
// using a service name suffixed by a dash and a random number.
func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
