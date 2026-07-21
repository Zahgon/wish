package wish

import (
	"context"
	"io"
	"os/exec"

	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/ssh"
)

func CommandContext(ctx context.Context, s ssh.Session, name string, args ...string) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

func Command(s ssh.Session, name string, args ...string) *Cmd {
	_ = "STUB: not implemented"
	return nil
}

type Cmd struct {
	sess ssh.Session
	cmd  *exec.Cmd
}

func (c *Cmd) SetEnv(env []string) { _ = "STUB: not implemented"; return }

func (c *Cmd) Environ() []string { _ = "STUB: not implemented"; return nil }

func (c *Cmd) SetDir(dir string) { _ = "STUB: not implemented"; return }

func (c *Cmd) Run() error { _ = "STUB: not implemented"; return nil }

var _ tea.ExecCommand = &Cmd{}

func (*Cmd) SetStderr(io.Writer) { _ = "STUB: not implemented"; return }

func (*Cmd) SetStdin(io.Reader) { _ = "STUB: not implemented"; return }

func (*Cmd) SetStdout(io.Writer) { _ = "STUB: not implemented"; return }
