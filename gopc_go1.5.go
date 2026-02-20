// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

//go:build (386 || amd64 || amd64p32 || arm || arm64 || s390x) && gc && go1.5

package gopc

import "unsafe"

const is32bit = ^uint(0)>>32 == 0

// Defined in gopc_go1.5.s.
func getg() unsafe.Pointer
