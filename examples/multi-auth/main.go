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
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"
	gossh "golang.org/x/crypto/ssh"
)

const (
	host          = "localhost"
	port          = "23234"
	validPassword = "asd123"
)

var users = map[string]string{
	"Carlos": "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILxWe2rXKoiO6W14LYPVfJKzRfJ1f3Jhzxrgjc/D4tU7",
}

func main() {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),

		wish.WithPublicKeyAuth(func(_ ssh.Context, key ssh.PublicKey) bool {
			log.Info("publickey")
			for _, pubkey := range users {
				parsed, _, _, _, _ := ssh.ParseAuthorizedKey(
					[]byte(pubkey),
				)
				if ssh.KeysEqual(key, parsed) {
					return true
				}
			}
			return false
		}),

		wish.WithPasswordAuth(func(_ ssh.Context, password string) bool {
			log.Info("password")
			return password == validPassword
		}),

		wish.WithKeyboardInteractiveAuth(func(_ ssh.Context, challenger gossh.KeyboardInteractiveChallenge) bool {
			log.Info("keyboard-interactive")
			answers, err := challenger(
				"", "",
				[]string{
					"♦ How much is 2+3: ",
					"♦ Which editor is best, vim or emacs? ",
				},
				[]bool{true, true},
			)
			if err != nil {
				return false
			}

			return len(answers) == 2 && answers[0] == "5" && answers[1] == "vim"
		}),

		wish.WithMiddleware(
			logging.Middleware(),
			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					wish.Println(sess, "Authorized!")
					wish.Println(sess, sess.PublicKey())
				}
			},
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
