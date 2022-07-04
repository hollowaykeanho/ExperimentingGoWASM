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

type SignalOSChannel adapterSignalOSChannel

// Known Signal Values from syscall package
const (
	SIGNAL_SIGABRT   = uint16(0x06)
	SIGNAL_SIGALRM   = uint16(0x0e)
	SIGNAL_SIGBUS    = uint16(0x07)
	SIGNAL_SIGCHLD   = uint16(0x11)
	SIGNAL_SIGCLD    = uint16(0x11)
	SIGNAL_SIGCONT   = uint16(0x12)
	SIGNAL_SIGFPE    = uint16(0x08)
	SIGNAL_SIGHUP    = uint16(0x01)
	SIGNAL_SIGILL    = uint16(0x04)
	SIGNAL_SIGINT    = uint16(0x02)
	SIGNAL_SIGIO     = uint16(0x1d)
	SIGNAL_SIGIOT    = uint16(0x06)
	SIGNAL_SIGKILL   = uint16(0x09)
	SIGNAL_SIGPIPE   = uint16(0x0d)
	SIGNAL_SIGPOLL   = uint16(0x1d)
	SIGNAL_SIGPROF   = uint16(0x1b)
	SIGNAL_SIGPWR    = uint16(0x1e)
	SIGNAL_SIGQUIT   = uint16(0x03)
	SIGNAL_SIGSEGV   = uint16(0x0b)
	SIGNAL_SIGSTKFLT = uint16(0x10)
	SIGNAL_SIGSTOP   = uint16(0x13)
	SIGNAL_SIGSYS    = uint16(0x1f)
	SIGNAL_SIGTERM   = uint16(0x0f)
	SIGNAL_SIGTRAP   = uint16(0x05)
	SIGNAL_SIGTSTP   = uint16(0x14)
	SIGNAL_SIGTTIN   = uint16(0x15)
	SIGNAL_SIGTTOU   = uint16(0x16)
	SIGNAL_SIGUNUSED = uint16(0x1f)
	SIGNAL_SIGURG    = uint16(0x17)
	SIGNAL_SIGUSR1   = uint16(0x0a)
	SIGNAL_SIGUSR2   = uint16(0x0c)
	SIGNAL_SIGVTALRM = uint16(0x1a)
	SIGNAL_SIGWINCH  = uint16(0x1c)
	SIGNAL_SIGXCPU   = uint16(0x18)
	SIGNAL_SIGXFSZ   = uint16(0x19)
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
// It accepts:
//   1. `bufferSize` = the size of the channel buffer (default: 3).
//
// It shall returns:
//   1. `hestiaError.OK` | `0` = Successful
//   2. `hestiaError.ENOENT` | `2` = given parameter is `nil`.
func SignalInit(sig *Signal, bufferSize int) hestiaError.Error {
	return _signalInit(sig, bufferSize)
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

// SignalSend sends a given signal into the hestiaOS.Signal object.
//
// It shall returns:
//   1. `hestiaError.OK` | `0` = Successful
//   2. `hestiaError.ENOENT` | `2` = given parameter is `nil`.
//   3. `hestiaError.` | `2` = given parameter is `nil`.
func SignalSend(sig *Signal, data uint16) hestiaError.Error {
	return _signalSend(sig, data)
}

// SignalSubscribeOS set a channel to receive signal from the OS.
func SignalSubscribeOS(ch SignalOSChannel) {
	_signalSubscribeOS(ch)
}

// SignalUnsubscribeOS set a channel to stop receive signal from the OS.
func SignalUnsubscribeOS(ch SignalOSChannel) {
	_signalUnsubscribeOS(ch)
}
