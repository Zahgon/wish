package scp

import (
	"io"
	"io/fs"

	"charm.land/wish/v2"
	"github.com/charmbracelet/ssh"
)

type CopyToClientHandler interface {
	Glob(ssh.Session, string) ([]string, error)

	WalkDir(ssh.Session, string, fs.WalkDirFunc) error

	NewDirEntry(ssh.Session, string) (*DirEntry, error)

	NewFileEntry(ssh.Session, string) (*FileEntry, func() error, error)
}

type CopyFromClientHandler interface {
	Mkdir(ssh.Session, *DirEntry) error

	Write(ssh.Session, *FileEntry) (int64, error)
}

type Handler interface {
	CopyFromClientHandler
	CopyToClientHandler
}

func Middleware(rh CopyToClientHandler, wh CopyFromClientHandler) wish.Middleware {
	_ = "STUB: not implemented"
	return *new(wish.Middleware)
}

var NULL = []byte{'\x00'}

type Entry interface {
	Write(io.Writer) error

	path() string
}

type AppendableEntry interface {
	Write(io.Writer) error

	Append(entry Entry)
}

type FileEntry struct {
	Name     string
	Filepath string
	Mode     fs.FileMode
	Size     int64
	Reader   io.Reader
	Atime    int64
	Mtime    int64
}

func (e *FileEntry) path() string { _ = "STUB: not implemented"; return "" }

func (e *FileEntry) Write(w io.Writer) error { _ = "STUB: not implemented"; return nil }

type RootEntry []Entry

func (e *RootEntry) Append(entry Entry) { _ = "STUB: not implemented"; return }

func (e *RootEntry) Write(w io.Writer) error { _ = "STUB: not implemented"; return nil }

//nolint:wrapcheck

type DirEntry struct {
	Children []Entry
	Name     string
	Filepath string
	Mode     fs.FileMode
	Atime    int64
	Mtime    int64
}

func (e *DirEntry) path() string { _ = "STUB: not implemented"; return "" }

func (e *DirEntry) Write(w io.Writer) error { _ = "STUB: not implemented"; return nil }

//nolint:wrapcheck

func (e *DirEntry) Append(entry Entry) { _ = "STUB: not implemented"; return }

type Op byte

const (
	OpCopyToClient Op = 'f'

	OpCopyFromClient Op = 't'
)

type Info struct {
	Ok bool

	Recursive bool

	Path string

	Op Op
}

func GetInfo(cmd []string) Info { _ = "STUB: not implemented"; return *new(Info) }

func octalPerms(info fs.FileMode) string { _ = "STUB: not implemented"; return "" }

func normalizePath(p string) string { _ = "STUB: not implemented"; return "" }

func validateName(name string) error { _ = "STUB: not implemented"; return nil }
