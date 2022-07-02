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
)

// App is the app activity life-cycle approach kernel structure.
//
// The design was based on iOS and Android mobile application life-cycle
// approach. This kernel structure focused on event-driven signalling for
// triggering based-level activities either on its own (OS-independent) or
// interacting with the opreating system signals like power management and
// interrupts.
//
// This kernel structure also respects operating system's signals like SIGINT,
// SIGTERM, and SIGKILL shall stop and destroy the executing app.
type App struct {
	// OnCreate is for initializing all resources ready to use.
	//
	// It accepts a `State` parameter which can be any data type that
	// contains the saved state of the App.
	//
	// This function is **OPTIONAL**.
	OnCreate func(any)

	// OnDestroy is for exiting the application.
	//
	// This function shall destroy the application, freeing its own
	// resources and conclude the app executions.
	//
	// This function is **OPTIONAL**.
	OnDestroy func()

	// OnPause is for handling for pausing or entering power sleep state.
	//
	// This function shall prepare the application for temporary acitvity
	// halting or power sleep state where its preparations is for later
	// restorations.
	OnPause func()

	// OnRestart is for any restart executions prior to a `Start` function.
	//
	// The function is designed solely for early interaction for a heavy
	// and long waiting `Start` function.
	//
	// This function is **OPTIONAL**.
	OnRestart func()

	// SignalRestart is the interrupt value for executing OnRestart().
	//
	// To enable such signal, it **SHALL NOT** be zero or conflicting with
	// all defined hestiaOS.SIGNAL values.
	SignalRestart uint16

	// OnResume is for returning back from paused or power slept state.
	//
	// This function shall restore the application back to operational
	// mode.
	OnResume func()

	// OnStart is for starting the Application.
	//
	// This function shall start the application executions.
	//
	// This function is **COMPULSORY**.
	OnStart func()

	// OnStop is for stopping the Application.
	//
	// This function shall stop the application but not destroying it. At
	// this state, it can be started again.
	//
	// This function is **COMPULSORY**.
	OnStop func()

	signaler *hestiaOS.Signal

	// ServerMode is to set the Application into a server behavior.
	//
	// Setting this to `true` shall stop it from emitting an automated
	// signal similar to `AppKill(...)`, allowing the appplication close
	// after the execution. This means you're entirely responsible for
	// terminating the app by calling `AppKill(...)` function.
	//
	// Default is `false` to behave like a common program.
	ServerMode bool
}

// AppKill signals a running app to stop running and destroy itself.
func AppKill(app *App) (err hestiaError.Error) {
	return _appKill(app)
}

// AppPause signals a running app to enter its pause state.
func AppPause(app *App) (err hestiaError.Error) {
	return _appPause(app)
}

// AppRestart signals a running app to restart itself.
func AppRestart(app *App) (err hestiaError.Error) {
	return _appRestart(app)
}

// AppResume signals a paused app to its running state.
func AppResume(app *App) (err hestiaError.Error) {
	return _appResume(app)
}

// AppRun executes a given app from create to destroy.
//
// It accepts 3 parameters:
//   1. `app` - the App Kernel Data Structure object
//   2. `state` - the input state object for `OnCreate(...)` function.
//   3. `buffer` - the size of the signal buffer. Default is 3.
//
// It returns:
//   1. `OK` | `0` - fully executed.
//   2. `ENOENT` | `2` - `app` parameter is missing.
//   3. `ENOPROTOOPT` | `92` - either `OnStop` or `OnStart` or both are missing.
//   4. all hestiaOS.SignalInit error codes.
func AppRun(app *App, state any, buffer int) (err hestiaError.Error) {
	if buffer == 0 {
		buffer = 3
	}

	return _appRun(app, state, buffer)
}

func appSignal(app *App, value uint16) (err hestiaError.Error) {
	err = appValidate(app)
	if err != hestiaError.OK {
		return err
	}

	if app.signaler == nil {
		return hestiaError.ENOPROTOOPT
	}

	if value == 0x00 {
		return hestiaError.ENOMSG
	}

	hestiaOS.SignalSend(app.signaler, value)

	return hestiaError.OK
}

func appValidate(app *App) (err hestiaError.Error) {
	if app == nil {
		return hestiaError.ENOENT
	}

	if app.OnStop == nil {
		return hestiaError.ENOPROTOOPT
	}

	if app.OnStart == nil {
		return hestiaError.ENOPROTOOPT
	}

	return hestiaError.OK
}
