// Copyright IBM Corp. 2018, 2020
// SPDX-License-Identifier: Apache-2.0

// +build appengine

package logrus

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
