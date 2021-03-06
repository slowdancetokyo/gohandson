// Copyright 2017 The CRT Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 amd64p32 arm64 mips64 mips64le mips64p32 mips64p32le ppc64 sparc64

package crt // import "modernc.org/ccgo/crt"

func VALong(ap *[]interface{}) int64   { return VAInt64(ap) }
func VAULong(ap *[]interface{}) uint64 { return VAUint64(ap) }
