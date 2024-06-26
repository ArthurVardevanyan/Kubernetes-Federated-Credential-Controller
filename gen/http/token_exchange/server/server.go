// Code generated by goa v3.16.1, DO NOT EDIT.
//
// tokenExchange HTTP server
//
// Command:
// $ goa gen k8s-federated-credential-api/design

package server

import (
	"context"
	tokenexchange "k8s-federated-credential-api/gen/token_exchange"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the tokenExchange service endpoint HTTP handlers.
type Server struct {
	Mounts        []*MountPoint
	ExchangeToken http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the tokenExchange service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *tokenexchange.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ExchangeToken", "POST", "/exchangeToken"},
		},
		ExchangeToken: NewExchangeTokenHandler(e.ExchangeToken, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "tokenExchange" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ExchangeToken = m(s.ExchangeToken)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return tokenexchange.MethodNames[:] }

// Mount configures the mux to serve the tokenExchange endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountExchangeTokenHandler(mux, h.ExchangeToken)
}

// Mount configures the mux to serve the tokenExchange endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountExchangeTokenHandler configures the mux to serve the "tokenExchange"
// service "exchangeToken" endpoint.
func MountExchangeTokenHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/exchangeToken", f)
}

// NewExchangeTokenHandler creates a HTTP handler which loads the HTTP request
// and calls the "tokenExchange" service "exchangeToken" endpoint.
func NewExchangeTokenHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeExchangeTokenRequest(mux, decoder)
		encodeResponse = EncodeExchangeTokenResponse(encoder)
		encodeError    = EncodeExchangeTokenError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "exchangeToken")
		ctx = context.WithValue(ctx, goa.ServiceKey, "tokenExchange")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}
