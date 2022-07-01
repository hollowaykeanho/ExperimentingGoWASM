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

//go:build wasm
// +build wasm

package hestiaKernel

import (
	"hestiaGo/hestiaError"
	"hestiaGo/hestiaOS"
)

func _appRun(app *App, state interface{}, buffer int) (err hestiaError.Error) {
	var signal uint16

	// validate all inputs
	err = appValidate(app)
	if err != hestiaError.OK {
		return err
	}

	// initialize key variables
	app.signaler = &hestiaOS.Signal{}
	err = hestiaOS.SignalInit(app.signaler, buffer)
	if err != hestiaError.OK {
		return err
	}

	// create the app environment first
	if app.OnCreate != nil {
		app.OnCreate(state)
	}

	// start the app
	app.OnStart()

	// wait for interrupt signal
	for {
		signal = hestiaOS.SignalWait(app.signaler)

		switch {
		case signal == 0x00:
			continue
		case signal == app.SignalRestart:
			if app.OnRestart != nil {
				app.OnRestart()
			}
			app.OnStop()
			app.OnStart()
		case signal == hestiaOS.SIGNAL_SIGSTOP:
			if app.OnPause != nil {
				app.OnPause()
			}
		case signal == hestiaOS.SIGNAL_SIGCONT:
			if app.OnResume != nil {
				app.OnResume()
			}
		case signal == hestiaOS.SIGNAL_SIGINT:
			goto end
		case signal == hestiaOS.SIGNAL_SIGTERM:
			goto end
		case signal == hestiaOS.SIGNAL_SIGKILL:
			goto end
		default:
		}
	}

end:
	app.OnStop()
	if app.OnDestroy != nil {
		app.OnDestroy()
	}

	return hestiaError.OK
}

func _appKill(app *App) (err hestiaError.Error) {
	return appSignal(app, hestiaOS.SIGNAL_SIGKILL)
}

func _appPause(app *App) (err hestiaError.Error) {
	return appSignal(app, hestiaOS.SIGNAL_SIGSTOP)
}

func _appRestart(app *App) (err hestiaError.Error) {
	return appSignal(app, app.SignalRestart)
}

func _appResume(app *App) (err hestiaError.Error) {
	return appSignal(app, hestiaOS.SIGNAL_SIGCONT)
}
