package bubbletea

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/wish/v2"
	"github.com/charmbracelet/ssh"
)

type BubbleTeaHandler = Handler //nolint:revive

type Handler func(sess ssh.Session) (tea.Model, []tea.ProgramOption)

type ProgramHandler func(sess ssh.Session) *tea.Program

func Middleware(handler Handler) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

func MiddlewareWithProgramHandler(handler ProgramHandler) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

func MakeOptions(sess ssh.Session) []tea.ProgramOption { _ = "STUB: not implemented"; return nil }

func newDefaultProgramHandler(handler Handler) ProgramHandler {
	_ = "STUB: not implemented"
	return *new(ProgramHandler)
}
