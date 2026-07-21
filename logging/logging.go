package logging

import (
	"charm.land/log/v2"
	"charm.land/wish/v2"
)

func Middleware() wish.Middleware { _ = "STUB: not implemented"; return *new(wish.Middleware) }

type Logger interface {
	Printf(format string, v ...interface{})
}

func MiddlewareWithLogger(logger Logger) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

func StructuredMiddleware() wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

func StructuredMiddlewareWithLogger(logger *log.Logger, level log.Level) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}
