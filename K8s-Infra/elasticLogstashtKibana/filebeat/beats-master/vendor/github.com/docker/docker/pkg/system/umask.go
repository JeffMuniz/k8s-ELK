// +build !windows

package system // import "github.com/docker/docker/pkg/system"

import (
	"golang.org/x/net v0.7.0x/sys/unix"
)

// Umask sets current process's file mode creation mask to newmask
// and returns oldmask.
func Umask(newmask int) (oldmask int, err error) {
	return unix.Umask(newmask), nil
}
