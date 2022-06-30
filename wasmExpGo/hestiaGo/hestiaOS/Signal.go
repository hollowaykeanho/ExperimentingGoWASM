// Copyright 2022 "Holloway" Chew, Kean Ho <hollowaykeanho@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package hestiaOS

import (
	"hestiaGo/hestiaError"
)

// Known Signal Values
const (
	SIGNAL_SIGINT  = uint16(0x02) // from syscall.SIGINT via os/signal
	SIGNAL_SIGKILL = uint16(0x09) // from syscall.SIGKILL via os/signal
	SIGNAL_SIGSTOP = uint16(0x13) // from syscall.SIGSTOP
	SIGNAL_SIGTERM = uint16(0x0f) // from syscall.SIGTERM
)

// Signal is for event-driven trigger independent on platform OS.
//
// Unlike syscall.Signal, hestiaOS's Signal permits the use of its own feature
// outside of platform operating system signals. This is very useful for
// operating in a non-OS environment like microcontroller and wasm.
//
// hestiaOS.Signal object **REQUIRES** initialization via `InitSignal(...)`
// function. Using it without initialization **SHALL YIELD UNPREDICTABLE
// CONCEQUENCES**. Hence, it's your duty to initialize it right after creation.
//
// To ensure there is enough signal ID, a 16-bit unsigned integer is used.
type Signal struct {
	channel chan uint16
}

// SignalInit initializes the hestiaOS.Signal object.
//
// It shall returns:
//   1. `hestiaError.OK` | `0` = Successful
//   2. `hestiaError.ENOENT` | `2` = given parameter is `nil`.
func SignalInit(sig *Signal) hestiaError.Error {
	return _signalInit(sig)
}

// SignalWait sets the hestiaOS.Signal object to wait for a signal.
//
// Depending on deployed platform or OS, WaitSignal do listen to operating
// system signal on top of user's own signal values.
//
// Therefore, you're advised to use values away from the operating system's
// signal values (e.g. `syscall.Signal` constant list).
func SignalWait(sig *Signal) uint16 {
	return _signalWait(sig)
}

// SignalStop sends SIGNAL_SIGSTOP into the hestiaOS.Signal object.
//
// It shall returns:
//   1. `hestiaError.OK` | `0` = Successful
//   2. `hestiaError.ENOENT` | `2` = given parameter is `nil`.
func SignalStop(sig *Signal) hestiaError.Error {
	return _signalStop(sig)
}

// SignalSend sends a given signal into the hestiaOS.Signal object.
//
// It shall returns:
//   1. `hestiaError.OK` | `0` = Successful
//   2. `hestiaError.ENOENT` | `2` = given parameter is `nil`.
//   3. `hestiaError.` | `2` = given parameter is `nil`.
func SignalSend(sig *Signal, data uint16) hestiaError.Error {
	return _signalSend(sig, data)
}
