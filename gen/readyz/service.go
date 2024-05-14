// Code generated by goa v3.16.1, DO NOT EDIT.
//
// readyz service
//
// Command:
// $ goa gen k8s-federated-credential-api/design

package readyz

import (
	"context"
)

// Service is the readyz service interface.
type Service interface {
	// Readyz implements readyz.
	Readyz(context.Context) (res *ReadyzResult, err error)
}

// APIName is the name of the API as defined in the design.
const APIName = "kfca"

// APIVersion is the version of the API as defined in the design.
const APIVersion = "0.0.1"

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "readyz"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"readyz"}

// ReadyzResult is the result type of the readyz service readyz method.
type ReadyzResult struct {
	Ready bool
}
