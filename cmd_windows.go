//go:build windows
// +build windows

package wish

import (
	"github.com/charmbracelet/ssh"
)

func (c *Cmd) doRun(ppty ssh.Pty, _ <-chan ssh.Window) error { _ = "STUB: not implemented"; return nil }

//nolint:wrapcheck
