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

//go:build !wasm
// +build !wasm

package hestiaOS

import (
	"hestiaGo/hestiaError"
	"os"
	"os/signal"
	"syscall"
)

func _signalInit(sig *Signal) hestiaError.Error {
	if sig == nil {
		return hestiaError.ENOENT
	}

	sig.channel = make(chan uint16, 1)

	return hestiaError.OK
}

func _signalSend(sig *Signal, data uint16) hestiaError.Error {
	if sig == nil {
		return hestiaError.ENOENT
	}

	sig.channel <- data

	return hestiaError.OK
}

func _signalStop(sig *Signal) hestiaError.Error {
	if sig == nil {
		return hestiaError.ENOENT
	}

	sig.channel <- SIGNAL_SIGSTOP

	return hestiaError.OK
}

func _signalWait(sig *Signal) (value uint16) {
	var ok bool
	var osSig os.Signal
	var sysSignal syscall.Signal
	var chOS chan os.Signal

	if sig == nil {
		goto done
	}

	// make a temporary OS-signal channel for its notices
	chOS = make(chan os.Signal, 1)
	signal.Notify(chOS, os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
		syscall.SIGSTOP,
	)

	for {
		select {
		case value, ok = <-sig.channel:
			if !ok {
				continue
			}

			signal.Stop(chOS)
			goto done
		case osSig, ok = <-chOS:
			if !ok {
				continue
			}

			sysSignal = osSig.(syscall.Signal)

			sig.channel <- uint16(sysSignal)
		default:
		}
	}

done:
	return value
}
