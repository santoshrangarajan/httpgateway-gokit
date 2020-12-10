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
	log.Println("Bookservice Isavailable:name", name)
	if name == "" {
		return false, ErrEmpty
	}
	log.Println("Bookservice Isavailable, returning true")
	return true, nil
}

func (bookService) Authorname(name string) (string, error) {
	if name == "" {
		return "", ErrEmpty
	}
	return "santosh", nil
}

func (bookService) Count(name string) int {
	log.Println("Bookservice count:name", name)
	if name == "" {
		return 0
	}
	return 100
}

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

	log.Println("starting....")
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
