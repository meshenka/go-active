package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type lastSeenRequest struct {
	Identifier string `json:"identifer"`
}

type lastSeenResponse struct {
	LastSeen string `json:"lastSeen"`
	Err      string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type updateRequest struct {
	Identifier string `json:"identifier"`
}

type updateResponse struct {
	Visitor jsonVisitor `json:"visitor"`
	Err     string      `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type jsonVisitor struct {
	Identifier string `json:"identifer"`
	LastSeen   string `json:"lastSeen"`
}

func makeLastSeenEndpoint(svc LastSeenService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(lastSeenRequest)
		time, err := svc.LastSeen(req.Identifier)

		if err != nil {
			return lastSeenResponse{time.String(), err.Error()}, nil
		}

		return lastSeenResponse{time.String(), ""}, nil
	}
}

func makeUpdateEndpoint(svc LastSeenService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(updateRequest)
		visitor, err := svc.Update(req.Identifier)

		var v = jsonVisitor{visitor.identifier, visitor.lastSeen.String()}
		if err != nil {
			return updateResponse{v, err.Error()}, nil
		}

		return updateResponse{v, ""}, nil
	}
}
