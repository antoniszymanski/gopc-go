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
	ch := make(chan string, 100)
	for range cap(ch) {
		go func() {
			pc := gopc.Get()
			frame, ok := FrameForPC(pc)
			switch {
			case !ok:
				ch <- "unable to get frame"
			case frame.Function != "github.com/antoniszymanski/gopc-go_test.Test":
				ch <- "unexpected function: " + frame.Function
			case path.Base(frame.File) != "gopc_test.go":
				ch <- "unexpected file: " + frame.File
			case frame.Line != 18:
				ch <- "unexpected line: " + strconv.Itoa(frame.Line)
			default:
				ch <- ""
			}
		}()
	}
	for range cap(ch) {
		if err := <-ch; err != "" {
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
