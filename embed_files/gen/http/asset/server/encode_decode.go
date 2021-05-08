// Code generated by goa v3.2.4, DO NOT EDIT.
//
// asset HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/embed_files/design

package server

import (
	"context"
	"net/http"
	"strconv"

	asset "goa.design/examples/embed_files/gen/asset"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodePublicResponse returns an encoder for responses returned by the asset
// public endpoint.
func EncodePublicResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*asset.PublicResult)
		val := res.Length
		lengths := strconv.FormatInt(val, 10)
		w.Header().Set("Content-Length", lengths)
		if res.ContentType != nil {
			w.Header().Set("Content-Type", *res.ContentType)
		}
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodePublicRequest returns a decoder for requests sent to the asset public
// endpoint.
func DecodePublicRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			filename string

			params = mux.Vars(r)
		)
		filename = params["filename"]
		payload := filename

		return payload, nil
	}
}

// EncodePublicError returns an encoder for errors returned by the public asset
// endpoint.
func EncodePublicError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "invalid_file_path":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewPublicInvalidFilePathResponseBody(res)
			}
			w.Header().Set("goa-error", "invalid_file_path")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewPublicInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}