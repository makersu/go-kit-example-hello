package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(HelloService) HelloService

type loggingMiddleware struct {
	logger log.Logger
	next   HelloService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a HelloService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next HelloService) HelloService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Hello(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Hello", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Hello(ctx, s)
}
