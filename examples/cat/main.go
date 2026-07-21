package main

import (
	"log"

	"charm.land/wish/v2"
	"charm.land/wish/v2/logging"
	"github.com/charmbracelet/ssh"
)

func middleware(sh ssh.Handler) ssh.Handler { _ = "STUB: not implemented"; return *new(ssh.Handler) }

func main() {
	s, err := wish.NewServer(

		ssh.AllocatePty(),
		wish.WithHostKeyPath("id_cat"),
		wish.WithAddress(":2022"),
		wish.WithMiddleware(
			logging.MiddlewareWithLogger(log.Default()),
			middleware,
		),
	)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	log.Printf("listening on %q", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
}
