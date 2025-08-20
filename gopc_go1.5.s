// SPDX-FileCopyrightText: 2021 Peter Mattis
// SPDX-License-Identifier: MPL-2.0

// Assembly to mimic runtime.getg.

//go:build (386 || amd64 || amd64p32 || arm || arm64 || s390x) && gc && go1.5
// +build 386 amd64 amd64p32 arm arm64 s390x
// +build gc
// +build go1.5

#include "textflag.h"

// func getg() unsafe.Pointer
TEXT ·getg(SB),NOSPLIT,$0-8
#ifdef GOARCH_386
	MOVL (TLS), AX
	MOVL AX, ret+0(FP)
#endif
#ifdef GOARCH_amd64
	MOVQ (TLS), AX
	MOVQ AX, ret+0(FP)
#endif
#ifdef GOARCH_arm
	MOVW g, ret+0(FP)
#endif
#ifdef GOARCH_arm64
	MOVD g, ret+0(FP)
#endif
#ifdef GOARCH_s390x
	MOVD g, ret+0(FP)
#endif
	RET
