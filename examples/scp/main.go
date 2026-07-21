package main

import (
	"context"
	"errors"
	"io"
	"io/fs"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"charm.land/log/v2"
	"charm.land/wish/v2"
	"charm.land/wish/v2/scp"
	"github.com/charmbracelet/ssh"
	"github.com/pkg/sftp"
)

const (
	host = "localhost"
	port = "23235"
)

func main() {
	root, _ := filepath.Abs("./examples/scp/testdata")
	handler := scp.NewFileSystemHandler(root)
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),

		wish.WithSubsystem("sftp", sftpSubsystem(root)),
		wish.WithMiddleware(

			scp.Middleware(handler, handler),
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

func sftpSubsystem(root string) ssh.SubsystemHandler {
	_ = "STUB: not implemented"
	return *new(ssh.SubsystemHandler)
}

type sftpHandler struct {
	root string
}

var (
	_ sftp.FileLister = &sftpHandler{}
	_ sftp.FileReader = &sftpHandler{}
)

type listerAt []fs.FileInfo

func (l listerAt) ListAt(ls []fs.FileInfo, offset int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (s *sftpHandler) Fileread(r *sftp.Request) (io.ReaderAt, error) {
	_ = "STUB: not implemented"
	return *new(io.ReaderAt), nil
}

func (s *sftpHandler) Filelist(r *sftp.Request) (sftp.ListerAt, error) {
	_ = "STUB: not implemented"
	return *new(sftp.ListerAt), nil
}
