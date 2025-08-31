// SPDX-FileCopyrightText: 2025 Antoni Szymański
// SPDX-License-Identifier: MPL-2.0

package gopc_test

import (
	"path"
	"runtime"
	"strconv"
	"testing"

	"github.com/antoniszymanski/gopc-go"
)

func Test(t *testing.T) {
	errs := make(chan string, 100)
	for range cap(errs) {
		go func() {
			pc := gopc.Get()
			frame, ok := FrameForPC(pc)
			switch {
			case !ok:
				errs <- "unable to get frame"
			case frame.Function != "github.com/antoniszymanski/gopc-go_test.Test":
				errs <- "unexpected function: " + frame.Function
			case path.Base(frame.File) != "gopc_test.go":
				errs <- "unexpected file: " + frame.File
			case frame.Line != 18:
				errs <- "unexpected line: " + strconv.Itoa(frame.Line)
			default:
				errs <- ""
			}
		}()
	}
	for range cap(errs) {
		if err := <-errs; err != "" {
			t.Fatal(err)
		}
	}
}

func FrameForPC(pc uintptr) (frame runtime.Frame, ok bool) {
	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ = frames.Next()
	return frame, frame.PC != 0
}

func Benchmark(b *testing.B) {
	for range b.N {
		gopc.Get()
	}
}
