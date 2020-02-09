package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := lastSeenService{}
	updateHandler := httptransport.NewServer(
		makeUpdateEndpoint(svc),
		decodeUpdateRequest,
		encodeResponse,
	)

	lastSeenHandler := httptransport.NewServer(
		makeLastSeenEndpoint(svc),
		decodeLastSeenRequest,
		encodeResponse,
	)

	http.Handle("/update", updateHandler)
	http.Handle("/get", lastSeenHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request updateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeLastSeenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request lastSeenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
