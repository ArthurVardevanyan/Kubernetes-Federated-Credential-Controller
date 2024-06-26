// Code generated by goa v3.16.1, DO NOT EDIT.
//
// tokenExchange endpoints
//
// Command:
// $ goa gen k8s-federated-credential-api/design

package tokenexchange

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "tokenExchange" service endpoints.
type Endpoints struct {
	ExchangeToken goa.Endpoint
}

// NewEndpoints wraps the methods of the "tokenExchange" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		ExchangeToken: NewExchangeTokenEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "tokenExchange" service
// endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ExchangeToken = m(e.ExchangeToken)
}

// NewExchangeTokenEndpoint returns an endpoint function that calls the method
// "exchangeToken" of service "tokenExchange".
func NewExchangeTokenEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ExchangeTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		ctx, err = authJWTFn(ctx, p.Authorization, &sc)
		if err != nil {
			return nil, err
		}
		return s.ExchangeToken(ctx, p)
	}
}
