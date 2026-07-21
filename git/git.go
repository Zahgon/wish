package git

import (
	"errors"

	"charm.land/wish/v2"
	"github.com/charmbracelet/ssh"
)

var ErrNotAuthed = errors.New("you are not authorized to do this")

var ErrSystemMalfunction = errors.New("something went wrong")

var ErrInvalidRepo = errors.New("invalid repo")

type AccessLevel int

const (
	NoAccess AccessLevel = iota

	ReadOnlyAccess

	ReadWriteAccess

	AdminAccess
)

type GitHooks = Hooks //nolint:revive

type Hooks interface {
	AuthRepo(string, ssh.PublicKey) AccessLevel
	Push(string, ssh.PublicKey)
	Fetch(string, ssh.PublicKey)
}

func Middleware(repoDir string, gh Hooks) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

//nolint:exhaustive

//nolint:exhaustive

func gitPack(s ssh.Session, gitCmd string, repoDir string, repo string) error {
	_ = "STUB: not implemented"
	return nil
}

func fileExists(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func Fatal(s ssh.Session, v ...interface{}) { _ = "STUB: not implemented"; return }

func EnsureRepo(dir, repo string) error { _ = "STUB: not implemented"; return nil }

func runGit(s ssh.Session, dir string, args ...string) error { _ = "STUB: not implemented"; return nil }

func ensureDefaultBranch(s ssh.Session, repoPath string) error {
	_ = "STUB: not implemented"
	return nil
}
