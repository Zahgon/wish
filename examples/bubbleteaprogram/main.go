package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/log/v2"
	"charm.land/wish/v2"
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"
)

const (
	host = "localhost"
	port = "23234"
)

func main() {
	s, err := wish.NewServer(
		ssh.AllocatePty(),
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			myCustomBubbleteaMiddleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		log.Error("Could not stop server", "error", err)
	}
}

func myCustomBubbleteaMiddleware() wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

type model struct {
	term   string
	width  int
	height int
	time   time.Time
}

type timeMsg time.Time

func (m model) Init() tea.Cmd { _ = "STUB: not implemented"; return *new(tea.Cmd) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

func (m model) View() tea.View { _ = "STUB: not implemented"; return *new(tea.View) }
