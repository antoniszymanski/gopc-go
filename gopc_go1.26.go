// SPDX-FileCopyrightText: 2026 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

//go:build gc && go1.26 && !go1.27

package gopc

import "unsafe"

func Get() uintptr {
	if is32bit {
		return *(*uintptr)(unsafe.Add(getg(), 184))
	} else {
		return *(*uintptr)(unsafe.Add(getg(), 288))
	}
}
