package recover

import (
	"charm.land/wish/v2"
)

func Middleware(mw ...wish.Middleware) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

type Logger interface {
	Printf(format string, v ...interface{})
}

func MiddlewareWithLogger(logger Logger, mw ...wish.Middleware) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}
