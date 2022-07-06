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

package main

import (
	"fmt"

	"hestiaGo/hestiaError"
	"hestiaGo/hestiaKernel/hestiaAppKernel"
	"hestiaGo/hestiaOS/hestiaWASM"
)

func onCreate() {
	fmt.Printf("Initializing wasmExpGo run...\n")

	// setup a simple promise
	promise := &hestiaWASM.Promise{
		Name: "myGoFx",
		Func: func() hestiaError.Error {
			fmt.Printf("from promised world\n")
			h2, _ := hestiaWASM.CreateElement("h2")

			html := []byte("Render from Promise!")
			_ = hestiaWASM.SetHTML(h2, &html)

			body := hestiaWASM.Body()
			_ = hestiaWASM.Append(body, h2)
			return hestiaError.OK
		},
		Resolve: func() any {
			fmt.Printf("promise resolved!\n")
			return "promise resolved!"
		},
		Reject: func(hestiaError.Error) any {
			fmt.Printf("promise rejected!\n")
			return "promise rejected!"
		},
	}
	hestiaWASM.GoPromise(promise)
}

func onStart() {
	// test hestiaWASM basic functions
	fmt.Printf("[ Track 1] Starting wasmExpGo run...\n")

	fmt.Printf("'69' Convertable? %#v\n",
		hestiaWASM.IsTypeConvertable(69) == hestiaError.OK)

	x, xerr := hestiaWASM.ExecJSFunc(true, "myGoFx")
	fmt.Printf("Exec Promise? %#v %#v\n", x, xerr)

	// test HTML I/O
	fmt.Printf("[ Track 2] create HTML button element...\n")
	go uiInit()
}

func onStop() {
	fmt.Printf("Stopping wasmExpGo run...\n")
}

func main() {
	app := &hestiaAppKernel.Kernel{}
	hestiaAppKernel.SetFunction(app, hestiaAppKernel.FUNCTION_CREATE, onCreate)
	hestiaAppKernel.SetFunction(app, hestiaAppKernel.FUNCTION_START, onStart)
	hestiaAppKernel.SetFunction(app, hestiaAppKernel.FUNCTION_STOP, onStop)
	hestiaAppKernel.SetServerMode(app, hestiaAppKernel.MODE_SERVER)

	hestiaAppKernel.Run(app, 3)
}
