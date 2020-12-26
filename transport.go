package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// For each method, we define request and response structs
type isavailableRequest struct {
	Name string `json:"name"`
}

type isavailableResponse struct {
	V   bool   `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

// For each method, we define request and response structs
type countRequest struct {
	Name string `json:"name"`
}

// For each method, we define request and response structs
type countResponse struct {
	Count int `json:"count"`
}

// For each method, we define request and response structs
type authornameRequest struct {
	Name string `json:"name"`
}

type authornameResponse struct {
	Name string `json:"name"`
	Err  string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func makeAuthornameEndpoint(bksvc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(authornameRequest)
		v, err := bksvc.Authorname(req.Name)
		if err != nil {
			return authornameResponse{v, err.Error()}, nil
		}
		return authornameResponse{v, ""}, nil
	}
}

func makeCountEndpoint(bksvc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		count := bksvc.Count(req.Name)
		return countResponse{count}, nil
	}
}

func makeIsavailableEndpoint(bksvc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(isavailableRequest)
		v, err := bksvc.Isavailable(req.Name)
		if err != nil {
			return isavailableResponse{false, err.Error()}, nil
		}
		log.Println("makeIsavailableEndpoint, returning response")
		return isavailableResponse{v, ""}, nil
	}
}

func decodeAuthornameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request authornameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeIsAvailableRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request isavailableRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
