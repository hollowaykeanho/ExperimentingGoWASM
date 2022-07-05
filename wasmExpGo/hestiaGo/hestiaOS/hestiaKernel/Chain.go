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

package hestiaKernel

import (
	"hestiaGo/hestiaError"
	"hestiaGo/hestiaOS"
	"sync"
)

// Chain is the kernel structure for chain-executing function blocks.
//
// This kernal approach orignated from chain design architecture where function
// execution blocks are executed and next block is determined on the fly. In
// other word, the previous block decide and setup the next function block. This
// is suitable for event-driven oriented executions such as but not limited to
// user-driven activities or interrupt-driven activities.
//
// Chain is designed to be thread-safe in order to work both stand-alone as a
// kernel in non-OS environment or on top of another kernel structure
// (e.g. App).
//
// This kernel structure relies heavily on `func(any) any` function pattern
// where the previous block's output value is the input of the next block. It
// **ONLY HANDLES ONE (1)** parameter so if you need multiple ones, consider
// defining a data `struct` or using a `map` data list. To avoid restricting
// the kernel to a certain data type, the chain shall accept `any` data type and
// it's up to the user to assert the data type inside the function block. One
// easy way is:
//       func myBlock(arg any) any {
//           out, ok := arg.(MyType)
//           if !ok {
//               return nil
//           }
//
//           ...
//       }
type Chain struct {
	signaler *hestiaOS.Signal
	next     func(any) any
	mutex    *sync.Mutex
}

// ChainHasNext checks a given chain has a next function block to execute.
//
// This function is thread safe and is designed for thread-safe checking
// purposes.
//
// It accepts the following parameters:
//   1. `kernel` - the Chain kernel object.
//
// It returns the following outputs:
//   1. hestiaError.OK | `0` - YES, there is a next function block.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is `nil`.
//   3. hestiaError.ENODATA | `61` - NO, the next block is `nil`.
func ChainHasNext(kernel *Chain) hestiaError.Error {
	var hasNext bool

	if kernel == nil {
		return hestiaError.EOWNERDEAD
	}
	_chainInitMutex(kernel)

	kernel.mutex.Lock()
	hasNext = (kernel.next != nil)
	kernel.mutex.Unlock()

	if hasNext {
		return hestiaError.OK
	}

	return hestiaError.ENODATA
}

// ChainSetNext chains a given function block to be the next execution block.
//
// This function is thread safe and is designed for thread-safe assigning
// purposes. It is best to assign the next block right before returning a value
// in the current block.
//
// Note that the current block's return value is the next block's input
// parameter.
//
// It accepts the following parameters:
//   1. `kernel` - the Chain kernel object.
//   1. `fx` - the next function block.
//
// It returns the following outputs:
//   1. hestiaError.OK | `0` - successful assignment.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is `nil`.
//   3. hestiaError.ENODATA | `61` - given function `f` is `nil`.
func ChainSetNext(kernel *Chain, fx func(any) any) (err hestiaError.Error) {
	if kernel == nil {
		return hestiaError.EOWNERDEAD
	}
	_chainInitMutex(kernel)

	if fx == nil {
		return hestiaError.ENODATA
	}

	kernel.mutex.Lock()
	kernel.next = fx
	kernel.mutex.Unlock()

	return hestiaError.OK
}

// ChainSignal triggers an chain execution event with a given signal.
//
// This function is thread safe and is designed for thread-safe execution
// purposes.
//
// It accepts the following parameters:
//   1. `kernel` - the Chain kernel object.
//   1. `signal` - a given signal.
//
// It returns the following outputs:
//   1. hestiaError.OK | `0` - signal sent successfully.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is `nil`.
//   3. hestiaError.EHOSTUNREACH | `61` - the chain signaler is `nil` (idling?)
func ChainSignal(kernel *Chain, signal uint16) (err hestiaError.Error) {
	if kernel == nil {
		return hestiaError.EOWNERDEAD
	}
	_chainInitMutex(kernel)

	if _chainIsSignalerNil(kernel) {
		return hestiaError.EHOSTUNREACH
	}

	hestiaOS.SignalSend(_chainGetSignaler(kernel), signal)

	return hestiaError.OK
}

// ChainStart initializes and run a given Chain kernel object.
//
// This function is thread safe and is designed for thread-safe execution
// purposes. However, **IT IS A BLOCKING FUNCTION** as it turns itself into a
// signal listening server. Hence, use `go ...` go routine whenever a
// non-blocking execution is required.
//
// This kernel listens to these special signals for stopping the server:
//     1. `hestiaOS.SIGNAL_SIGSTOP`
//     2. `hestiaOS.SIGNAL_SIGINT`
//     3. `hestiaOS.SIGNAL_SIGTERM`
//     4. `hestiaOS.SIGNAL_SIGKILL`
// If the kernel is operating in an OS environment, the OS signals can influence
// the kernel directly. Should any of these signals are given to the first
// function block, it is required to perform graceful shutdown in this ONE (1)
// block itself as **the next block SHALL NOT be executed at all**.
//
// This kernel listen to these special signals and other unidentified signal for
// triggering a chain event:
//     1. `hestiaOS.SIGNAL_SIGCONT`
//
// This function accepts the first function block (`first`) that only accepts
// hestiaOS.Signal (`uint16`) stated above value as its input parameter. The
// goal of the first block is to distinguish next block based on the triggered
// action and also formulating the required return value for the chain.
//
// It accepts the following parameters:
//   1. `kernel` - the Chain kernel object.
//   2. `first` - the first function block.
//   3. `buffer` - the signal buffer size. Default is `3`.
//
// It returns the following outputs:
//   1. hestiaError.OK | `0` - successful server execution and termination.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is `nil`.
//   3. hestiaError.ENODATA | `61` - given function `first` is `nil`.
//   4. hestiaError.ESTRPIPE | `86` - failed to initialize signaler.
func ChainStart(kernel *Chain, first func(any) any, buffer int) (err hestiaError.Error) {
	var ret any
	var signal uint16
	var function func(any) any

	if kernel == nil {
		return hestiaError.EOWNERDEAD
	}
	_chainInitMutex(kernel)

	if buffer == 0 {
		buffer = 3
	}

	if first == nil {
		return hestiaError.ENODATA
	}
	ChainSetNext(kernel, first)

	if _chainIsSignalerNil(kernel) {
		kernel.signaler = &hestiaOS.Signal{}
	}

	if hestiaOS.SignalInit(_chainGetSignaler(kernel), buffer) != hestiaError.OK {
		return hestiaError.ESTRPIPE
	}

listen:
	signal = hestiaOS.SignalWait(_chainGetSignaler(kernel))
	ret = signal

	// execute chain
	ChainSetNext(kernel, first)
	for ChainHasNext(kernel) == hestiaError.OK {
		function = _chainGetNext(kernel)
		ret = function(ret)
	}

	// reset progression back to first block
	function = nil
	ret = nil

	// decide next action based on signal
	switch signal {
	case hestiaOS.SIGNAL_SIGSTOP:
		goto end
	case hestiaOS.SIGNAL_SIGCONT:
		goto listen
	case hestiaOS.SIGNAL_SIGINT:
		goto end
	case hestiaOS.SIGNAL_SIGTERM:
		goto end
	case hestiaOS.SIGNAL_SIGKILL:
		goto end
	default:
		goto listen
	}

end:
	return hestiaError.OK
}

// ChainStop stops a running Chain object and set it to idle.
//
// It is an alias of:
//       ChainSignal(kernel, hestiaOS.SIGNAL_SIGSTOP)
func ChainStop(kernel *Chain) (err hestiaError.Error) {
	return ChainSignal(kernel, hestiaOS.SIGNAL_SIGSTOP)
}

// ChainTrigger triggers a listening Chain object for an event.
//
// It is an alias of:
//       ChainSignal(kernel, hestiaOS.SIGNAL_SIGCONT)
func ChainTrigger(kernel *Chain) (err hestiaError.Error) {
	return ChainSignal(kernel, hestiaOS.SIGNAL_SIGCONT)
}

func _chainGetNext(kernel *Chain) (out func(any) any) {
	kernel.mutex.Lock()
	out = kernel.next
	kernel.next = nil
	kernel.mutex.Unlock()

	return out
}

func _chainGetSignaler(kernel *Chain) (out *hestiaOS.Signal) {
	kernel.mutex.Lock()
	out = kernel.signaler
	kernel.mutex.Unlock()

	return out
}

func _chainInitMutex(kernel *Chain) {
	if kernel.mutex == nil {
		kernel.mutex = &sync.Mutex{}
	}
}

func _chainIsSignalerNil(kernel *Chain) (verdict bool) {
	kernel.mutex.Lock()
	verdict = (kernel.signaler == nil)
	kernel.mutex.Unlock()

	return verdict
}
