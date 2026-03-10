// Copyright IBM Corp. 2018, 2020
// SPDX-License-Identifier: Apache-2.0

// +build !appengine,!js,!windows,!nacl,!plan9

package logrus

import (
	"io"
	"os"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return isTerminal(int(v.Fd()))
	default:
		return false
	}
}
