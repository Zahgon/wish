package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"charm.land/log/v2"
	"charm.land/wish/v2"
	"charm.land/wish/v2/git"
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"
)

const (
	port    = "23233"
	host    = "localhost"
	repoDir = ".repos"
)

type app struct {
	access git.AccessLevel
}

func (a app) AuthRepo(string, ssh.PublicKey) git.AccessLevel {
	_ = "STUB: not implemented"
	return *new(git.AccessLevel)
}

func (a app) Push(repo string, _ ssh.PublicKey) { _ = "STUB: not implemented"; return }

func (a app) Fetch(repo string, _ ssh.PublicKey) { _ = "STUB: not implemented"; return }

func main() {

	a := app{git.ReadWriteAccess}

	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),

		ssh.PublicKeyAuth(func(ssh.Context, ssh.PublicKey) bool { return true }),

		ssh.PasswordAuth(func(ssh.Context, string) bool { return false }),
		wish.WithMiddleware(

			git.Middleware(repoDir, a),

			gitListMiddleware,
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

func gitListMiddleware(next ssh.Handler) ssh.Handler {
	_ = "STUB: not implemented"
	return *new(ssh.Handler)
}
