package scp

import (
	"bufio"
	"regexp"

	"github.com/charmbracelet/ssh"
)

var (
	reTimestamp = regexp.MustCompile(`^T(\d{10}) 0 (\d{10}) 0$`)
	reNewFolder = regexp.MustCompile(`^D(\d{4}) 0 (.*)$`)
	reNewFile   = regexp.MustCompile(`^C(\d{4}) (\d+) (.*)$`)
)

type parseError struct {
	subject string
}

func (e parseError) Error() string { _ = "STUB: not implemented"; return "" }

func copyFromClient(s ssh.Session, info Info, handler CopyFromClientHandler) error {
	_ = "STUB: not implemented"
	return nil
}

func handleNewFile(s ssh.Session, r *bufio.Reader, handler CopyFromClientHandler, path, line string, match []string, mtime, atime int64) error {
	_ = "STUB: not implemented"
	return nil
}
