// Copyright IBM Corp. 2018, 2020
// SPDX-License-Identifier: Apache-2.0

// +build js nacl plan9

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
