// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build dragonfly freebsd netbsd openbsd solaris

package osfs

import (
	"golang.org/x/net v0.7.0x/sys/unix"

	"github.com/elastic/go-txfile/internal/vfs"
)

type syncState struct{}

// Sync uses fsync, for flushing and syncing a file to disk.  If the OS, file
// system, or disk drivers do not enforce a flush on all the intermediate
// caches and the drive itself, there is a chance of data loss and file
// corruption on power failure.
func (f *File) Sync(flags vfs.SyncFlag) error {
	// best effort
	for {
		err := f.File.Sync()
		if err == nil || (err != unix.EINTR && err != unix.EAGAIN) {
			return f.wrapErr("file/sync", err)
		}
	}
}
