package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

////https://github.com/tensor-programming/go-kit-tutorial
////https://sagikazarmark.hu/blog/getting-started-with-go-kit/

// BookService is ...
type BookService interface {
	Isavailable(string) (bool, error)
	Authorname(string) (string, error)
	Count(string) int
}

type bookService struct{}

// ErrEmpty is ...
var ErrEmpty = errors.New("empty string")

func (bookService) Isavailable(name string) (bool, error) {
	if name == "" {
		return false, ErrEmpty
	}
	return true, nil
}

func (bookService) Authorname(name string) (string, error) {
	if name == "" {
		return "", ErrEmpty
	}
	return "santosh", nil
}

func (bookService) Count(s string) int {
	if s == "" {
		return 0
	}
	return 0
}

// For each method, we define request and response structs
type isavailableRequest struct {
	N string `json:"name"`
}

type isavailableResponse struct {
	V   bool   `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

// For each method, we define request and response structs
type countRequest struct {
	N string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

// For each method, we define request and response structs
type authornameRequest struct {
	N string `json:"name"`
}

type authornameResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

func makeAuthornameEndpoint(bksvc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(authornameRequest)
		v, err := bksvc.Authorname(req.N)
		if err != nil {
			return authornameResponse{v, err.Error()}, nil
		}
		return authornameResponse{v, ""}, nil
	}
}

func makeCountEndpoint(bksvc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := bksvc.Count(req.N)
		return countResponse{v}, nil
	}
}

func makeIsavailableEndpoint(bksvc BookService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(isavailableRequest)
		v, err := bksvc.Isavailable(req.N)
		if err != nil {
			return isavailableResponse{false, err.Error()}, nil
		}
		return isavailableResponse{v, err.Error()}, nil
	}
}

// Endpoints

func main() {
	bksvc := bookService{}

	authornameHandler := httptransport.NewServer(
		makeAuthornameEndpoint(bksvc),
		decodeAuthornameRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(bksvc),
		decodeCountRequest,
		encodeResponse,
	)

	isavailableHandler := httptransport.NewServer(
		makeIsavailableEndpoint(bksvc),
		decodeIsAvailableRequest,
		encodeResponse,
	)

	http.Handle("/authorname", authornameHandler)
	http.Handle("/count", countHandler)
	http.Handle("/isavailable", isavailableHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	log.Println("Go-Kit POC")
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

/////// end point
/////// transport
/////// encapsulation of objects inside service
