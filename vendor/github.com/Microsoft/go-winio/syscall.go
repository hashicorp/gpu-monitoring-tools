// Copyright IBM Corp. 2018, 2020
// SPDX-License-Identifier: Apache-2.0

package winio

//go:generate go run $GOROOT/src/syscall/mksyscall_windows.go -output zsyscall_windows.go file.go pipe.go sd.go fileinfo.go privilege.go backup.go hvsock.go
