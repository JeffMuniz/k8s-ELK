// +build darwin dragonfly freebsd netbsd openbsd

package terminal

import "golang.org/x/net v0.7.0x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA

func IsTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}
