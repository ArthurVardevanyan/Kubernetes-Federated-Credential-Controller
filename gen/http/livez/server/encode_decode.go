// Code generated by goa v3.16.1, DO NOT EDIT.
//
// livez HTTP server encoders and decoders
//
// Command:
// $ goa gen k8s-federated-credential-api/design

package server

import (
	"context"
	livez "k8s-federated-credential-api/gen/livez"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeLivezResponse returns an encoder for responses returned by the livez
// livez endpoint.
func EncodeLivezResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*livez.LivezResult)
		enc := encoder(ctx, w)
		body := NewLivezResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}
