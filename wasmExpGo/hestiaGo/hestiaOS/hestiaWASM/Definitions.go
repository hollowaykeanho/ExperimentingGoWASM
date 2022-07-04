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

// EventPhase is the W3C DOM Event Flow representations.
//
// More info: https://www.w3.org/TR/DOM-Level-3-Events/#event-flow
type EventPhase uint8

// EventPhase representations ID
const (
	EVENT_PHASE_CAPTURE  EventPhase = 1
	EVENT_PHASE_TARGET   EventPhase = 2
	EVENT_PHASE_BUBBLING EventPhase = 3
)

// Body() returns the DOM `<body>` Object got from Document().
func Body() *Object {
	return _body()
}

// Document() returns the DOM Document Object got from Global.
func Document() *Object {
	return _document()
}

// Event is the returning event return value called by Javascript Event.
//
// These data were translated back to Go Format for adaptive compliance.
type Event struct {
	// IsBubble states the current event trigger dispatch mode.
	//
	// This checks against bubble mode which is the default behavior.
	IsBubble bool

	// IsCancellable states the capability of cancelling this event.
	IsCancellable bool

	// IsComposed states the even capable of bubbling between shadow DOM.
	//
	// `true` means the event can bubble between shadow DOM and regular DOM.
	IsComposed bool

	// CurrentTarget refers to this event's currently registered target.
	CurrentTarget *Object

	// DefaultPrevented indicates the `preventDefault()` was called to this event.
	//
	// `true` means this event was called after the `event.preventDefault()`
	// API was called.
	DefaultPrevented bool

	// Phase indicates the current event phase.
	//
	// Refers to EventPhase constants list for each identity.
	Phase EventPhase

	// IsTrusted indicates the event was initiated by the browser or script.
	//
	// This is to distinguish the event call using script like event
	// creation method.
	IsTrusted bool

	// Target refers to this event's orignally dispatched.
	Target *Object

	// This refers to the current object.
	This *Object

	// Timestamp in milliseconds of the event.
	//
	// Refer to browser's specification for definitions.
	Timestamp float64

	// Type is the case-insensitive name.
	Type string
}

// EventListener is the adapter data structure for JS.addEventListener.
//
// The purpose is to serve as a standard approach from Go to Javascript
// instruction for adding/removing event listener into a Javascript object.
//
// For keeping this adapter package sane, **ONLY ONE (1)** EventListener can
// **ONLY BE ADDED TO ONE (1) ACTIVE OWNER AT A TIME**. If you need to serve
// multiple common listeners with the same function, have every common
// EventListener.Functions calls a common Go function instead. The sole reason
// is to preserve the handler for `RemoveEventListener()` at will.
//
// That also being said that it is ultimately your responsibility to keep the
// EventListener object for memory management.
type EventListener struct {
	// Name of the event.
	Name string

	// The listenting function to execute.
	//
	// This function is wrapped under hestiaWASM protecting function in a
	// separate goroutine.
	Function func(*Event)

	// CaptureMode is the decision for "bubble" or "capture" dispatch modes.
	//
	// "capture" mode is where the event of the outer parent is triggered
	// first then towards the deepest child element.
	//
	// "bubble" mode is where the event of the deepest child is triggered
	// first then towards the outest parent element.
	//
	// Default (`false`) is "bubble".
	Capture bool

	// Once is the decision for invoking the Function 1 time.
	//
	// As stated in Javascript standards, if Once is set to `true`, the
	// function shall be triggered 1-time and **SHALL** be removed
	// automatically.
	//
	// Default (`false`) is "always / not once".
	Once bool

	// Passive is the decision for operating the Function passively.
	//
	// This also means that the generated Javascript listening function
	// **SHALL NOT** opreate active listening functions like
	// `preventDefault()` and console warning **SHALL** be logged by the
	// Javascript execution counterpart.
	//
	// This is set to `true` automatically for real-time events like
	// `wheel`, `mousewheel`, `touchstart`, `touchmove`, and etc. Consult
	// Javascript specification for more info.
	Passive bool

	// PreventDefault is the decision for preventing default executions.
	//
	// Default (`false`) is allow default executions. If set to `true`,
	// event.PreventDefault() shall be invoked as soon as possible, even
	// before hestiaWASM.Event Go format parsing.
	PreventDefault bool

	// Signal is the abort signal use to trigger an abort handling.
	//
	// This is for inputting a general abort signal action into the
	// event listener.
	// TODO: **Requires Abort investigation needs**

	handler *Object
}

// Global() returns the DOM global Object.
func Global() *Object {
	return _global()
}

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
