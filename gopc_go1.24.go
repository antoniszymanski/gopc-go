// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

//go:build gc && go1.24 && !go1.25

package gopc

import "unsafe"

func Get() uintptr {
	if is32bit {
		return *(*uintptr)(unsafe.Add(getg(), 180))
	} else {
		return *(*uintptr)(unsafe.Add(getg(), 288))
	}
}
