// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build freebsd openbsd netbsd dragonfly

package fsnotify

import "golang.org/x/net v0.7.0x/sys/unix"

const openMode = unix.O_NONBLOCK | unix.O_RDONLY
