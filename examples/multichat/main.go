package main

import (
	"charm.land/bubbles/v2/textarea"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/ssh"
)

const (
	host = "localhost"
	port = "23234"
)

type app struct {
	*ssh.Server
	progs []*tea.Program
}

func (a *app) send(msg tea.Msg) { _ = "STUB: not implemented"; return }

func newApp() *app { _ = "STUB: not implemented"; return nil }

func (a *app) Start() { _ = "STUB: not implemented"; return }

func (a *app) ProgramHandler(s ssh.Session) *tea.Program { _ = "STUB: not implemented"; return nil }

func main() {
	app := newApp()
	app.Start()
}

type (
	errMsg  error
	chatMsg struct {
		id   string
		text string
	}
)

type model struct {
	*app
	viewport    viewport.Model
	messages    []string
	id          string
	textarea    *textarea.Model
	senderStyle lipgloss.Style
	err         error
}

func initialModel() model { _ = "STUB: not implemented"; return *new(model) }

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }
