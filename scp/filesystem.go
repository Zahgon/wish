package scp

import (
	"io/fs"

	"github.com/charmbracelet/ssh"
)

type fileSystemHandler struct{ root string }

var _ Handler = &fileSystemHandler{}

func NewFileSystemHandler(root string) Handler { _ = "STUB: not implemented"; return *new(Handler) }

func (h *fileSystemHandler) chtimes(path string, mtime, atime int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *fileSystemHandler) prefixed(path string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (h *fileSystemHandler) Glob(_ ssh.Session, s string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:wrapcheck

//nolint:wrapcheck

func (h *fileSystemHandler) WalkDir(_ ssh.Session, path string, fn fs.WalkDirFunc) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:wrapcheck

func (h *fileSystemHandler) NewDirEntry(_ ssh.Session, name string) (*DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *fileSystemHandler) NewFileEntry(_ ssh.Session, name string) (*FileEntry, func() error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (h *fileSystemHandler) Mkdir(_ ssh.Session, entry *DirEntry) error {
	_ = "STUB: not implemented"
	return nil
}

func (h *fileSystemHandler) Write(_ ssh.Session, entry *FileEntry) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

//nolint:errcheck
