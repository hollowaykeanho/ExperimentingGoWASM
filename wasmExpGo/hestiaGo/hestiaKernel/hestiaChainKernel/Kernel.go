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

package hestiaChainKernel

import (
	"hestiaGo/hestiaError"
	"hestiaGo/hestiaOS"
	"sync"
)

// Kernel is the Chain data structure.
type Kernel struct {
	signaler *hestiaOS.Signal
	next     func(any) any
	mutex    *sync.Mutex
}

// HasNext checks a given kernel having a next function block to execute.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// It accepts the following parameters:
//   1. `kernel` - the Kernel object.
//
// It returns the following outputs:
//   1. hestiaError.OK | `0` - there is a next function block.
//   2. hestiaError.ENODATA | `61` - there is NO next block (`nil`).
//   3. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                       `Validate(...)` function to diagnose.
func HasNext(kernel *Kernel) (err hestiaError.Error) {
	var verdict bool

	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	kernel.mutex.Lock()
	verdict = (kernel.next != nil)
	kernel.mutex.Unlock()

	if verdict {
		return hestiaError.OK
	}

	return hestiaError.ENODATA
}

// IsRunning checks a given kernel's running status.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - the given `kernel` is running.
//   2. hestiaError.EHOSTDOWN | `112` - the given `kernel` is idling.
//   3. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                       `Validate(...)` function to diagnose.
func IsRunning(kernel *Kernel) (err hestiaError.Error) {
	var verdict bool

	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	kernel.mutex.Lock()
	verdict = (kernel.signaler != nil)
	kernel.mutex.Unlock()

	if verdict {
		return hestiaError.OK
	}

	return hestiaError.EHOSTDOWN
}

// SetNext chains a given function block to be the next execution block.
//
// It is best to assign the next block right before returning a value in the
// current block. Note that the current block's return value is the next block's
// input parameter.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// It accepts the following parameters:
//   1. `kernel` - the Kernel object.
//   1. `fx` - the next function block.
//
// It returns the following outputs:
//   1. hestiaError.OK | `0` - successful assignment.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                       `Validate(...)` function to diagnose.
//   3. hestiaError.ENODATA | `61` - given function `f` is `nil`.
func SetNext(kernel *Kernel, fx func(any) any) (err hestiaError.Error) {
	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	if fx == nil {
		return hestiaError.ENODATA
	}

	kernel.mutex.Lock()
	kernel.next = fx
	kernel.mutex.Unlock()

	return hestiaError.OK
}

// Signal sends a given signal into the kernel for an event.
//
// This signal shall be fed as input into the first function block.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// It accepts the following parameters:
//   1. `kernel` - the Kernel object.
//   1. `signal` - a given signal.
//
// It returns the following outputs:
//   1. hestiaError.OK | `0` - signal sent successfully.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                       `Validate(...)` function to diagnose.
//   3. hestiaError.EHOSTUNREACH | `61` - the chain signaler is idling.
func Signal(kernel *Kernel, signal uint16) (err hestiaError.Error) {
	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	err = IsRunning(kernel)
	if err != hestiaError.OK {
		return hestiaError.EHOSTUNREACH
	}

	hestiaOS.SignalSend(_getSignaler(kernel), signal)

	return hestiaError.OK
}

// Start initializes and execute a given Kernel object.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// Due to it turning itself into a signal listening server,
// **IT IS A BLOCKING FUNCTION**. Hence, use `go ...` go routine when a
// non-blocking execution is required.
//
// This kernel listens to any of these special signals for stopping the server:
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
//   1. `kernel` - the Kernel object.
//   2. `first` - the first function block.
//   3. `buffer` - the signal buffer size. Default is `3`.
//
// It returns the following outputs:
//   1. hestiaError.OK | `0` - successful server execution and termination.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                       `Validate(...)` function to diagnose.
//   3. hestiaError.ENODATA | `61` - given function `first` is `nil`.
//   4. hestiaError.ESTRPIPE | `86` - failed to initialize signaler.
//   5. hestiaError.EBUSY | `16` - the kernel is already running.
func Start(kernel *Kernel, first func(any) any, buffer int) (err hestiaError.Error) {
	var ret any
	var signal uint16
	var function func(any) any

	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	err = IsRunning(kernel)
	if err == hestiaError.OK {
		return hestiaError.EBUSY
	}

	if buffer == 0 {
		buffer = 3
	}

	if first == nil {
		return hestiaError.ENODATA
	}
	SetNext(kernel, first)

	kernel.mutex.Lock()
	kernel.signaler = &hestiaOS.Signal{}
	err = hestiaOS.SignalInit(kernel.signaler, buffer)
	kernel.mutex.Unlock()

	if err != hestiaError.OK {
		return hestiaError.ESTRPIPE
	}

listen:
	signal = hestiaOS.SignalWait(_getSignaler(kernel))
	ret = signal
	function = nil

	// execute chain
	SetNext(kernel, first)
	for HasNext(kernel) == hestiaError.OK {
		function = _getNext(kernel)
		ret = function(ret)
	}
	ret = nil

	// decide next action based on signal
	switch signal {
	case hestiaOS.SIGNAL_SIGSTOP:
		goto end
	case hestiaOS.SIGNAL_SIGINT:
		goto end
	case hestiaOS.SIGNAL_SIGTERM:
		goto end
	case hestiaOS.SIGNAL_SIGKILL:
		goto end
	case hestiaOS.SIGNAL_SIGCONT:
		fallthrough
	default:
		goto listen
	}

end:
	return hestiaError.OK
}

// Stop sends a stop signal into a running Kernel object and set it to idle.
//
// It is an alias of:
//       Signal(kernel, hestiaOS.SIGNAL_SIGSTOP)
func Stop(kernel *Kernel) (err hestiaError.Error) {
	return Signal(kernel, hestiaOS.SIGNAL_SIGSTOP)
}

// Trigger sends an event signal into a running Kernel object for an event run.
//
// It is an alias of:
//       Signal(kernel, hestiaOS.SIGNAL_SIGCONT)
func Trigger(kernel *Kernel) (err hestiaError.Error) {
	return Signal(kernel, hestiaOS.SIGNAL_SIGCONT)
}

// Validate checks a given kernel is ready for operation.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following paramters:
//   1. `kernel` - the Kernel Object
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - given `kernel` is operable.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is `nil`.
func Validate(kernel *Kernel) (err hestiaError.Error) {
	if kernel == nil {
		return hestiaError.EOWNERDEAD
	}
	_configureMutex(kernel)

	return hestiaError.OK
}

func _configureMutex(kernel *Kernel) {
	if kernel.mutex == nil {
		kernel.mutex = &sync.Mutex{}
	}
}

func _getNext(kernel *Kernel) (out func(any) any) {
	kernel.mutex.Lock()
	out = kernel.next
	kernel.next = nil
	kernel.mutex.Unlock()

	return out
}

func _getSignaler(kernel *Kernel) (out *hestiaOS.Signal) {
	kernel.mutex.Lock()
	out = kernel.signaler
	kernel.mutex.Unlock()

	return out
}
