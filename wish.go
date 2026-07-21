package wish

import (
	"github.com/charmbracelet/ssh"
)

type Middleware func(next ssh.Handler) ssh.Handler

func NewServer(ops ...ssh.Option) (*ssh.Server, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:wrapcheck

//nolint:wrapcheck

//nolint:wrapcheck

func Fatal(s ssh.Session, v ...interface{}) { _ = "STUB: not implemented"; return }

func Fatalf(s ssh.Session, f string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Fatalln(s ssh.Session, v ...interface{}) { _ = "STUB: not implemented"; return }

func Error(s ssh.Session, v ...interface{}) { _ = "STUB: not implemented"; return }

func Errorf(s ssh.Session, f string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Errorln(s ssh.Session, v ...interface{}) { _ = "STUB: not implemented"; return }

func Print(s ssh.Session, v ...interface{}) { _ = "STUB: not implemented"; return }

func Printf(s ssh.Session, f string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Println(s ssh.Session, v ...interface{}) { _ = "STUB: not implemented"; return }

func WriteString(s ssh.Session, v string) (int, error) {
	_ = "STUB: not implemented"
	return 0,
		//nolint:wrapcheck
		nil
}
