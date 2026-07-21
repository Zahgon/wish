package main

import (
	"errors"
	"net"

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
	srv, err := wish.NewServer(

		wish.WithAddress(net.JoinHostPort(host, port)),

		wish.WithHostKeyPath(".ssh/id_ed25519"),

		wish.WithMiddleware(
			func(next ssh.Handler) ssh.Handler {
				return func(sess ssh.Session) {
					wish.Println(sess, "Hello, world!")
					next(sess)
				}
			},

			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	log.Info("Starting SSH server", "host", host, "port", port)
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {

		log.Error("Could not start server", "error", err)
	}
}
