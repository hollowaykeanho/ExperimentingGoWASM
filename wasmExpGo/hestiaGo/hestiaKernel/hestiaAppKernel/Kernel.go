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

package hestiaAppKernel

import (
	"hestiaGo/hestiaError"
	"hestiaGo/hestiaOS"
	"sync"
)

// FunctionType are the ID for selecting a function from the Kernel
type FunctionType int

const (
	// FUNCTION_CREATE is for initializing all resources ready to use.
	//
	// This function shall initializes the application, prepare its one-time
	// foundation for the `OnStart` function.
	//
	// This function is **OPTIONAL**.
	FUNCTION_CREATE FunctionType = iota + 1

	// FUNCTION_DESTROY is for graceful exiting the application.
	//
	// This function shall destroy the application, freeing its own
	// resources and conclude the app executions.
	//
	// This function is **OPTIONAL**.
	FUNCTION_DESTROY

	// FUNCTION_PAUSE is for handling pausing or entering sleep state.
	//
	// This function shall prepare the application for temporary acitvity
	// halting or power sleep state where its preparations is for later
	// restorations.
	//
	// Its output shall be restorable by FUNCTION_RESUME.
	//
	// This function is **OPTIONAL**.
	FUNCTION_PAUSE

	// FUNCTION_RESTART is for pre-restart executions before stop-start.
	//
	// The function is designed solely for user interaction for a heavy
	// and long waiting `Stop` and `Start` restarting function. It shall
	// run before the FUNCTION_STOP and then FUNCTION_START sequences.
	//
	// This function is **OPTIONAL**.
	FUNCTION_RESTART

	// FUNCTION_RESUME is for handling unpausing or exiting sleep state.
	//
	// This function shall restore the application back to operational
	// mode and unfreeze its operations from FUNCTION_PAUSE.
	//
	// This function is **OPTIONAL**.
	FUNCTION_RESUME

	// FUNCTION_START is for starting the kernel execution.
	//
	// This function shall start the executions of the application.
	//
	// This function is **COMPULSORY**.
	FUNCTION_START

	// FUNCTION_STOP is for stopping the Kernel executions.
	//
	// This function shall stop the application but not destroying it. At
	// this state, it can be started again via FUNCTION_RESTART.
	//
	// This function is **COMPULSORY**.
	FUNCTION_STOP
)

// RunMode are the ID for setting the Kernel operating mode.
type RunMode int

const (
	// MODE_ONCE is to set the kernel opreating once.
	//
	// This is the default behavior where `Kill(...)` is called
	// automatically right after the completion of FUNCTION_START execution.
	MODE_ONCE RunMode = iota

	// MODE_SERVER is to set the kernel operating as an active server.
	//
	// Due to the nature of the server mode, it is a blocking function.
	// Hence, please use `go Run(...)` goroutine for non-block requirements.
	//
	// To stop the server, simply kill the server via `Kill(...)` function.
	MODE_SERVER
)

// Kernel is the app activity life-cycle approach kernel structure.
//
// This kernel structure respects operating system's signals like SIGINT,
// SIGTERM, and SIGKILL shall stop and destroy the executing app.
type Kernel struct {
	onCreate  func()
	onDestroy func()
	onPause   func()
	onRestart func()
	onResume  func()
	onStart   func()
	onStop    func()

	signaler *hestiaOS.Signal
	mutex    *sync.Mutex

	restartSignal uint16
	mode          RunMode
}

// GetFunction is to obtain a kernel function from the Kernel object.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//   2. `fxType` - the function type value. See `FUNCTION_[ACTION]` constants
//                 list.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - fully executed.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                     `Validate(...)` function to diagnose.
//   3. hestiaError.EPROTONOSUPPORT | `93` - given `fxType` is unsupported.
func GetFunction(kernel *Kernel, fxType FunctionType) (f func(), err hestiaError.Error) {
	err = Validate(kernel)
	if err != hestiaError.OK {
		return nil, hestiaError.EOWNERDEAD
	}

	return getFunction(kernel, fxType)
}

func getFunction(kernel *Kernel, fxType FunctionType) (f func(), err hestiaError.Error) {
	switch fxType {
	case FUNCTION_CREATE:
		kernel.mutex.Lock()
		f = kernel.onCreate
		kernel.mutex.Unlock()
	case FUNCTION_DESTROY:
		kernel.mutex.Lock()
		f = kernel.onDestroy
		kernel.mutex.Unlock()
	case FUNCTION_PAUSE:
		kernel.mutex.Lock()
		f = kernel.onPause
		kernel.mutex.Unlock()
	case FUNCTION_RESTART:
		kernel.mutex.Lock()
		f = kernel.onRestart
		kernel.mutex.Unlock()
	case FUNCTION_RESUME:
		kernel.mutex.Lock()
		f = kernel.onResume
		kernel.mutex.Unlock()
	case FUNCTION_START:
		kernel.mutex.Lock()
		f = kernel.onStart
		kernel.mutex.Unlock()
	case FUNCTION_STOP:
		kernel.mutex.Lock()
		f = kernel.onStop
		kernel.mutex.Unlock()
	default:
		return nil, hestiaError.EPROTONOSUPPORT
	}

	return f, hestiaError.OK
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
//                                     `Validate(...)` function to diagnose.
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

// Kill sends a SIGNAL_SIGKILL signal to the kernel to stop its run.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - signal sent successfully.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                     `Validate(...)` function to diagnose.
//   3. hestiaError.EHOSTUNREACH | `113` - the given kernel is not running.
func Kill(kernel *Kernel) (err hestiaError.Error) {
	return signal(kernel, hestiaOS.SIGNAL_SIGKILL)
}

// Pause send a SIGNAL_SIGSTOP signal to the kernel executing FUNCTION_PAUSE.
//
// If the FUNCTION_PAUSE function was unset before running, then this signal
// shall perform nothing.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - signal sent successfully.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                     `Validate(...)` function to diagnose.
//   3. hestiaError.EHOSTUNREACH | `113` - the given kernel is not running.
func Pause(kernel *Kernel) (err hestiaError.Error) {
	return signal(kernel, hestiaOS.SIGNAL_SIGSTOP)
}

// Restart send restart signal to the kernel for restarting its executions.
//
// If the restart signal was unset or `0x00` using `SetSignal(...)` function
// before running, then this signal shall perform nothing.
//
// It shall trigger the execution in the following sequences:
//   1. FUNCTION_RESTART if available
//   2. FUNCTION_STOP
//   3. FUNCTION_START
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - signal sent successfully.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                     `Validate(...)` function to diagnose.
//   3. hestiaError.EHOSTUNREACH | `113` - the given kernel is not running.
func Restart(kernel *Kernel) (err hestiaError.Error) {
	var sig uint16

	kernel.mutex.Lock()
	sig = kernel.restartSignal
	kernel.mutex.Unlock()

	return signal(kernel, sig)
}

// Resume send a SIGNAL_SIGCONT signal to the kernel executing FUNCTION_RESUME.
//
// If the FUNCTION_RESUME function was unset before running, then this signal
// shall perform nothing.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - signal sent successfully.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                     `Validate(...)` function to diagnose.
//   3. hestiaError.EHOSTUNREACH | `113` - the given kernel is not running.
func Resume(kernel *Kernel) (err hestiaError.Error) {
	return signal(kernel, hestiaOS.SIGNAL_SIGCONT)
}

// Run executes a given kernel starting from create to destroy.
//
// As it shall turns itself into a signal listening server, this function is
// **A BLOCKING FUNCTION**. Please use `go hestiaAppKernel.Run(...)` goroutine
// for a non-blocking requirement.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel Data Structure object
//   2. `buffer` - the receiving signal buffer size. Default is 3.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - fully executed.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                     `Validate(...)` function to diagnose.
//   3. hestiaError.ESTRPIPE | `86` - failed to setup hestiaOS.Signal
func Run(kernel *Kernel, buffer int) (err hestiaError.Error) {
	var restartSig uint16
	var f func()

	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	if buffer == 0 {
		buffer = 3
	}

	// initialize signaler
	kernel.mutex.Lock()
	kernel.signaler = &hestiaOS.Signal{}
	err = hestiaOS.SignalInit(kernel.signaler, buffer)
	kernel.mutex.Unlock()

	if err != hestiaError.OK {
		return hestiaError.ESTRPIPE
	}

	// create app environment
	f, _ = getFunction(kernel, FUNCTION_CREATE)
	if f != nil {
		f()
	}

	// start the app
	f, _ = getFunction(kernel, FUNCTION_START)
	f()

	// signal non-server mode if unset
	if kernel.mode != MODE_SERVER {
		Kill(kernel)
	}

	// get restart signal
	kernel.mutex.Lock()
	restartSig = kernel.restartSignal
	kernel.mutex.Unlock()

	// listen to signal
listen:
	switch hestiaOS.SignalWait(_getSignaler(kernel)) {
	case 0x00:
		goto listen
	case restartSig:
		f, _ = getFunction(kernel, FUNCTION_RESTART)
		if f != nil {
			f()
		}

		f, _ = getFunction(kernel, FUNCTION_STOP)
		f()

		f, _ = getFunction(kernel, FUNCTION_START)
		f()

		goto listen
	case hestiaOS.SIGNAL_SIGSTOP:
		f, _ = getFunction(kernel, FUNCTION_PAUSE)
		if f != nil {
			f()
		}

		goto listen
	case hestiaOS.SIGNAL_SIGCONT:
		f, _ = getFunction(kernel, FUNCTION_RESUME)
		if f != nil {
			f()
		}

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
	f, _ = getFunction(kernel, FUNCTION_STOP)
	f()

	f, _ = getFunction(kernel, FUNCTION_DESTROY)
	if f != nil {
		f()
	}

	kernel.mutex.Lock()
	kernel.signaler = nil
	kernel.mutex.Unlock()

	return hestiaError.OK
}

// SetFunction is to configure a kernel function into the Kernel.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//   2. `fxType` - the function type value. See `FUNCTION_[ACTION]` constants
//                 list.
//   3. `f`      - the given function. `nil` is accepted for resetting the
//                 value.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - fully executed.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                     `Validate(...)` function to diagnose.
//   3. hestiaError.EPROTONOSUPPORT | `93` - given `fxType` is unsupported.
//   4. hestiaError.EBUSY | `16` - given kernel is currently running.
func SetFunction(kernel *Kernel, fxType FunctionType, f func()) (err hestiaError.Error) {
	err = Validate(kernel)
	if err == hestiaError.EOWNERDEAD {
		return hestiaError.EOWNERDEAD
	}

	err = IsRunning(kernel)
	if err == hestiaError.OK {
		return hestiaError.EBUSY
	}

	switch fxType {
	case FUNCTION_CREATE:
		kernel.mutex.Lock()
		kernel.onCreate = f
		kernel.mutex.Unlock()
	case FUNCTION_DESTROY:
		kernel.mutex.Lock()
		kernel.onDestroy = f
		kernel.mutex.Unlock()
	case FUNCTION_PAUSE:
		kernel.mutex.Lock()
		kernel.onPause = f
		kernel.mutex.Unlock()
	case FUNCTION_RESTART:
		kernel.mutex.Lock()
		kernel.onRestart = f
		kernel.mutex.Unlock()
	case FUNCTION_RESUME:
		kernel.mutex.Lock()
		kernel.onResume = f
		kernel.mutex.Unlock()
	case FUNCTION_START:
		kernel.mutex.Lock()
		kernel.onStart = f
		kernel.mutex.Unlock()
	case FUNCTION_STOP:
		kernel.mutex.Lock()
		kernel.onStop = f
		kernel.mutex.Unlock()
	default:
		return hestiaError.EPROTONOSUPPORT
	}

	return hestiaError.OK
}

// SetRestartSignal configures the kernel restart signal.
//
// By default, the signal is disabled (`0x00`). To enable such signal, simply
// assign any values not conflicting with any `hestiaOS.Signal` list is suffice.
// To reset the signal, simply set it back to `0x00`.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//   3. `signal` - the signal value. Default is disabled (`0x00`)
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - fully executed.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                     `Validate(...)` function to diagnose.
//   4. hestiaError.EBUSY | `16` - given kernel is currently running.
func SetRestartSignal(kernel *Kernel, signal uint16) (err hestiaError.Error) {
	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	err = IsRunning(kernel)
	if err == hestiaError.OK {
		return hestiaError.EBUSY
	}

	kernel.mutex.Lock()
	kernel.restartSignal = signal
	kernel.mutex.Unlock()

	return hestiaError.OK
}

// SetServerMode is to set the kernel operating mode.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//   2. `mode` - the run mode type value. See `MODE_[ACTION]` constants list.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - fully executed.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is unsable. Use
//                                       `Validate(...)` function to diagnose.
//   3. hestiaError.EPROTONOSUPPORT | `93` - given `mode` is unsupported.
//   4. hestiaError.EBUSY | `16` - given kernel is currently running.
func SetServerMode(kernel *Kernel, mode RunMode) (err hestiaError.Error) {
	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	err = IsRunning(kernel)
	if err == hestiaError.OK {
		return hestiaError.EBUSY
	}

	switch mode {
	case MODE_SERVER:
		kernel.mutex.Lock()
		kernel.mode = MODE_SERVER
		kernel.mutex.Unlock()
	case MODE_ONCE:
		kernel.mutex.Lock()
		kernel.mode = MODE_ONCE
		kernel.mutex.Unlock()
	default:
		return hestiaError.EPROTONOSUPPORT
	}

	return hestiaError.OK
}

func signal(kernel *Kernel, value uint16) (err hestiaError.Error) {
	err = Validate(kernel)
	if err != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	err = IsRunning(kernel)
	if err != hestiaError.OK {
		return hestiaError.EHOSTUNREACH
	}

	hestiaOS.SignalSend(_getSignaler(kernel), value)

	return hestiaError.OK
}

// Validate checks a given kernel is ready for operations.
//
// This function was designed to be operating the kernel in a thread-safe
// manner.
//
// This function accepts the following parameters:
//   1. `kernel` - the Kernel object.
//
// It returns any of the following output:
//   1. hestiaError.OK | `0` - given `kernel` is operable.
//   2. hestiaError.EOWNERDEAD | `130` - given `kernel` is `nil`.
//   3. hestiaError.ENOPROTOOPT | `92` - the kernel's `FUNCTION_START` and/or
//                                       `FUNCTION_STOP` are unset (`nil`).
func Validate(kernel *Kernel) (err hestiaError.Error) {
	var f func()

	if kernel == nil {
		return hestiaError.EOWNERDEAD
	}
	_configureMutex(kernel)

	// check onStop() function is set
	kernel.mutex.Lock()
	f = kernel.onStop
	kernel.mutex.Unlock()

	if f == nil {
		return hestiaError.ENOPROTOOPT
	}

	// check onStart() function is set
	kernel.mutex.Lock()
	f = kernel.onStart
	kernel.mutex.Unlock()

	if f == nil {
		return hestiaError.ENOPROTOOPT
	}

	return hestiaError.OK
}

func _configureMutex(kernel *Kernel) {
	if kernel.mutex == nil {
		kernel.mutex = &sync.Mutex{}
	}
}

func _getSignaler(kernel *Kernel) (out *hestiaOS.Signal) {
	kernel.mutex.Lock()
	out = kernel.signaler
	kernel.mutex.Unlock()

	return out
}
