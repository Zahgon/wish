package scp

import (
	"io/fs"

	"github.com/charmbracelet/ssh"
)

type fsHandler struct{ fsys fs.FS }

var _ CopyToClientHandler = &fsHandler{}

func NewFSReadHandler(fsys fs.FS) CopyToClientHandler {
	_ = "STUB: not implemented"
	return *new(CopyToClientHandler)
}

func (h *fsHandler) Glob(_ ssh.Session, s string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil,
		//nolint:wrapcheck
		nil
}

func (h *fsHandler) WalkDir(_ ssh.Session, path string, fn fs.WalkDirFunc) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:wrapcheck

func (h *fsHandler) NewDirEntry(_ ssh.Session, path string) (*DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fsHandler) NewFileEntry(_ ssh.Session, path string) (*FileEntry, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}
