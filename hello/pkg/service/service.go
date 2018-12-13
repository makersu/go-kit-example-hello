package service

import (
	"context"
	"errors"
)

// HelloService describes the service.
type HelloService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	Hello(ctx context.Context, s string) (rs string, err error)
}

type basicHelloService struct{}

// TODO implement the business logic of Hello
func (b *basicHelloService) Hello(ctx context.Context, s string) (rs string, err error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return "Hello, " + s, nil
}

// NewBasicHelloService returns a naive, stateless implementation of HelloService.
func NewBasicHelloService() HelloService {
	return &basicHelloService{}
}

// New returns a HelloService with all of the expected middleware wired in.
func New(middleware []Middleware) HelloService {
	var svc HelloService = NewBasicHelloService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
