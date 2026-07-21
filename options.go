package wish

import (
	"time"

	"github.com/charmbracelet/ssh"
)

func WithAddress(addr string) ssh.Option { _ = "STUB: not implemented"; return *new(ssh.Option) }

func WithVersion(version string) ssh.Option { _ = "STUB: not implemented"; return *new(ssh.Option) }

func WithBanner(banner string) ssh.Option { _ = "STUB: not implemented"; return *new(ssh.Option) }

func WithBannerHandler(h ssh.BannerHandler) ssh.Option {
	_ = "STUB: not implemented"
	return *new(ssh.Option)
}

func WithMiddleware(mw ...Middleware) ssh.Option {
	_ = "STUB: not implemented"
	return *new(ssh.Option)
}

func WithHostKeyPath(path string) ssh.Option { _ = "STUB: not implemented"; return *new(ssh.Option) }

func WithHostKeyPEM(pem []byte) ssh.Option { _ = "STUB: not implemented"; return *new(ssh.Option) }

func WithAuthorizedKeys(path string) ssh.Option { _ = "STUB: not implemented"; return *new(ssh.Option) }

func WithTrustedUserCAKeys(path string) ssh.Option {
	_ = "STUB: not implemented"
	return *new(ssh.Option)
}

func isAuthorized(path string, checker func(k ssh.PublicKey) bool) bool {
	_ = "STUB: not implemented"
	return false
}

//nolint:errcheck

func WithPublicKeyAuth(h ssh.PublicKeyHandler) ssh.Option {
	_ = "STUB: not implemented"
	return *new(ssh.Option)
}

func WithPasswordAuth(p ssh.PasswordHandler) ssh.Option {
	_ = "STUB: not implemented"
	return *new(ssh.Option)
}

func WithKeyboardInteractiveAuth(h ssh.KeyboardInteractiveHandler) ssh.Option {
	_ = "STUB: not implemented"
	return *new(ssh.Option)
}

func WithIdleTimeout(d time.Duration) ssh.Option {
	_ = "STUB: not implemented"
	return *new(ssh.Option)
}

func WithMaxTimeout(d time.Duration) ssh.Option { _ = "STUB: not implemented"; return *new(ssh.Option) }

func WithSubsystem(key string, h ssh.SubsystemHandler) ssh.Option {
	_ = "STUB: not implemented"
	return *new(ssh.Option)
}
