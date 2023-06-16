// +build linux aix

package terminal

import "golang.org/x/net v0.7.0x/sys/unix"

const ioctlReadTermios = unix.TCGETS

func IsTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}
