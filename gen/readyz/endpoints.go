// Code generated by goa v3.16.1, DO NOT EDIT.
//
// readyz endpoints
//
// Command:
// $ goa gen k8s-federated-credential-api/design

package readyz

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "readyz" service endpoints.
type Endpoints struct {
	Readyz goa.Endpoint
}

// NewEndpoints wraps the methods of the "readyz" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Readyz: NewReadyzEndpoint(s),
	}
}

// Use applies the given middleware to all the "readyz" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Readyz = m(e.Readyz)
}

// NewReadyzEndpoint returns an endpoint function that calls the method
// "readyz" of service "readyz".
func NewReadyzEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		return s.Readyz(ctx)
	}
}
