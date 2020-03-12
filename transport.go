package active

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// LastSeenRequest is a json structure of a call
type LastSeenRequest struct {
	Identifier string `json:"identifer"`
}

type lastSeenResponse struct {
	LastSeen string `json:"lastSeen"`
	Err      string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

// UpdateRequest json structure
type UpdateRequest struct {
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

// MakeLastSeenEndpoint build end point
func MakeLastSeenEndpoint(svc LastSeenService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(LastSeenRequest)
		time, err := svc.LastSeen(req.Identifier)

		if err != nil {
			return lastSeenResponse{time.String(), err.Error()}, nil
		}

		return lastSeenResponse{time.String(), ""}, nil
	}
}

// MakeUpdateEndpoint build endpoint
func MakeUpdateEndpoint(svc LastSeenService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		visitor, err := svc.Update(req.Identifier)

		var v = jsonVisitor{visitor.identifier, visitor.lastSeen.String()}
		if err != nil {
			return updateResponse{v, err.Error()}, nil
		}

		return updateResponse{v, ""}, nil
	}
}
