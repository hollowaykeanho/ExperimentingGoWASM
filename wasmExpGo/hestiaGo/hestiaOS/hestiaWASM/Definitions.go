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

package hestiaWASM

import (
	"hestiaGo/hestiaError"
)

// Object is the hestiaWASM adapter to syscall/js.Value object.
//
// The purpose is to ensure most of hestiaWASM is built on a stable environment
// while waiting for "syscall/js" to stabilize its own development.
type Object adapter

// Promise is the hestiaWASM adapter for Javscript Promise object.
//
// The goal is to become an adapter to Javascript Promise functionality for
// Go to fully utilize its own functionalities and syntax in its domain without
// getting the Javascript one mixed up. All the functional elements are basic
// execution values to keep the adapter abstraction as lean as possible.
//
// If you need specialized Go return values or similar behavior, please build it
// on top of these function values' layer.
//
// Do note that this Promise object is a stub that does nothing on a non-WASM
// platform.
type Promise struct {
	// Name is the promise function name used in Javascript domain.
	//
	// This field **SHALL NOT** be empty as the function can be called
	// externally via Javascript.
	//
	// This field also **SHALL NOT** contain characters incompatible with
	// Javascript function naming conventions.
	Name string

	// Func is the intended opreation to run inside a promise in pure Go.
	//
	// This outer function shall run with a goroutine (`go`) inside the
	// Javascript's Promise functionality for complying to its non-blocking
	// requirement.
	//
	// It **MUST AND SHALL** return a hestiaError.Error as its return value.
	// The following shall cause the Promise behaves accordingly after its
	// executions:
	//   1. hestiaError.OK | `0`    = Good ending: Invoke `Resolve(...)`
	//   2. hestiaError.(Any Error) = Bad ending: Invoke `Reject(...)`
	Func func() hestiaError.Error

	// Resolve is the invoke function opreate after a successful run.
	//
	// The return value shall be used for the Javascript Promise's
	// `Resolve.Invoke(...)` input. Hence, all return values **SHALL** be
	// Javascript compatible list of data stated in `syscall/js` package.
	Resolve func() any

	// Reject is the invoke function opreate after a failed run.
	//
	// The return value shall be used for the Javascript Promise's
	// `Reject.Invoke(...)` input. Hence, all return values **SHALL** be
	// Javascript compatible list of data stated in `syscall/js` package.
	//
	// This function takes 1 error code input where it indicates the return
	// status of Javascript due to possible alterations. The following
	// denotes the phenomenon:
	//   1. hestiaError.ENOTRECOVERABLE | `131` - corrupted JS callback
	//                                            functions (bad args)
	//   2. hestiaError.OK | `0` - Reject was properly invoked.
	Reject func(hestiaError.Error) any

	object *Object
}

// Global() returns the DOM global Object.
func Global() *Object {
	return _global()
}

// Document() returns the DOM Document Object got from Global.
func Document() *Object {
	return _document()
}

// Body() returns the DOM `<body>` Object got from Document().
func Body() *Object {
	return _body()
}
