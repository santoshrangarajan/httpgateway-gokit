package main

import (
	"errors"
	"log"
)

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
