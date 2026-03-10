// Copyright IBM Corp. 2018, 2020
// SPDX-License-Identifier: Apache-2.0

// +build darwin dragonfly freebsd netbsd openbsd

package logrus

import "golang.org/x/sys/unix"

const ioctlReadTermios = unix.TIOCGETA

func isTerminal(fd int) bool {
	_, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	return err == nil
}

