// Code generated by goa v3.2.4, DO NOT EDIT.
//
// asset client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/embed_files/design

package client

import (
	"context"
	"net/http"

	asset "goa.design/examples/embed_files/gen/asset"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the asset service endpoint HTTP clients.
type Client struct {
	// Public Doer is the HTTP client used to make requests to the public endpoint.
	PublicDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the asset service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		PublicDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Public returns an endpoint that makes HTTP requests to the asset service
// public server.
func (c *Client) Public() goa.Endpoint {
	var (
		decodeResponse = DecodePublicResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPublicRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PublicDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("asset", "public", err)
		}
		res, err := decodeResponse(resp)
		if err != nil {
			resp.Body.Close()
			return nil, err
		}
		return &asset.PublicResponseData{Result: res.(*asset.PublicResult), Body: resp.Body}, nil
	}
}