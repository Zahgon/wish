package ratelimiter

import (
	"errors"

	"charm.land/wish/v2"
	"github.com/charmbracelet/ssh"
	lru "github.com/hashicorp/golang-lru/v2"
	"golang.org/x/time/rate"
)

var ErrRateLimitExceeded = errors.New("rate limit exceeded, please try again later")

type RateLimiter interface {
	Allow(s ssh.Session) error
}

func Middleware(limiter RateLimiter) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

func NewRateLimiter(r rate.Limit, burst int, maxEntries int) RateLimiter {
	_ = "STUB: not implemented"
	return *new(RateLimiter)
}

type limiters struct {
	cache *lru.Cache[string, *rate.Limiter]
	rate  rate.Limit
	burst int
}

func (r *limiters) Allow(s ssh.Session) error { _ = "STUB: not implemented"; return nil }
