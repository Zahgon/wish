//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package bubbletea

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/ssh"
)

func makeOpts(s ssh.Session) []tea.ProgramOption { _ = "STUB: not implemented"; return nil }

//nolint:godox
