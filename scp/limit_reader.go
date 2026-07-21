package scp

import (
	"io"
	"sync"
)

func newLimitReader(r io.Reader, limit int) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

type limitReader struct {
	r io.Reader

	lock sync.Mutex
	left int
}

func (r *limitReader) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:wrapcheck
