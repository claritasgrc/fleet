package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/kolide/fleet/server/kolide"
	//"errors"
)

////////////////////////////////////////////////////////////////////////////////
// Test creating an endpoint
////////////////////////////////////////////////////////////////////////////////

type printRequest struct {
	Name string `json:"name"`
	Input string `json:"text"`
}

type printResponse struct {
	Output string `json:"response"`
	Err    error `json:"error,omitempty"`
}

func printInputEndpoint(svc kolide.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(printRequest)
		op, err := svc.PrintInputString(ctx, req.Name, req.Input)

		if err != nil {
			return printResponse{Err: err}, nil
		}
		return printResponse{Output: op}, nil
	}
}
