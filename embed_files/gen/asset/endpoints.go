// Code generated by goa v3.2.4, DO NOT EDIT.
//
// asset endpoints
//
// Command:
// $ goa gen goa.design/examples/embed_files/design

package asset

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "asset" service endpoints.
type Endpoints struct {
	Public goa.Endpoint
}

// PublicResponseData holds both the result and the HTTP response body reader
// of the "public" method.
type PublicResponseData struct {
	// Result is the method result.
	Result *PublicResult
	// Body streams the HTTP response body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "asset" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Public: NewPublicEndpoint(s),
	}
}

// Use applies the given middleware to all the "asset" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Public = m(e.Public)
}

// NewPublicEndpoint returns an endpoint function that calls the method
// "public" of service "asset".
func NewPublicEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(string)
		res, body, err := s.Public(ctx, p)
		if err != nil {
			return nil, err
		}
		return &PublicResponseData{Result: res, Body: body}, nil
	}
}
