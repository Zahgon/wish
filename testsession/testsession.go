package testsession

import (
	"net"
	"testing"

	"github.com/charmbracelet/ssh"
	gossh "golang.org/x/crypto/ssh"
)

func New(tb testing.TB, srv *ssh.Server, cfg *gossh.ClientConfig) *gossh.Session {
	_ = "STUB: not implemented"
	return nil
}

func Listen(tb testing.TB, srv *ssh.Server) string { _ = "STUB: not implemented"; return "" }

func newLocalListener(tb testing.TB) net.Listener {
	_ = "STUB: not implemented"
	return *new(net.Listener)
}

func NewClientSession(tb testing.TB, addr string, config *gossh.ClientConfig) (*gossh.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:gosec

//nolint:wrapcheck

//nolint:wrapcheck
