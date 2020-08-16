// Code generated by goa v3.2.3, DO NOT EDIT.
//
// calc gRPC client types
//
// Command:
// $ goa gen goa.design/examples/basic/design -o
// $(GOPATH)/src/goa.design/examples/basic

package client

import (
	calc "goa.design/examples/basic/gen/calc"
	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
)

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "calc" service.
func NewAddRequest(payload *calc.AddPayload) *calcpb.AddRequest {
	message := &calcpb.AddRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewAddResult builds the result type of the "add" endpoint of the "calc"
// service from the gRPC response type.
func NewAddResult(message *calcpb.AddResponse) int {
	result := int(message.Field)
	return result
}
