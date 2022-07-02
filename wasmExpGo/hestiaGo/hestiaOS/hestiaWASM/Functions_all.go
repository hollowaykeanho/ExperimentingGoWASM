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

package hestiaWASM

import (
	"hestiaGo/hestiaError"
)

// NOTE:
// This package is only meant for CPU=wasm or target=wasm build. To ensure
// import compatibility on other architecture, all functions SHALL be as
// follows:
//   1. output == unsupported { return hestiaError.EPFNOSUPPORT }
//   2. output == missing { return `nil` object }

func _appendChild(parent *Object, child *Object) hestiaError.Error {
	return hestiaError.EPROTONOSUPPORT
}

func _createElement(name string) (child *Object, err hestiaError.Error) {
	return nil, hestiaError.EPROTONOSUPPORT
}

func _get(parent *Object, query string) *Object {
	return nil
}

func _goPromise(promise *Promise) hestiaError.Error {
	return hestiaError.EPROTONOSUPPORT
}

func _isObjectOK(element *Object) hestiaError.Error {
	return hestiaError.EPROTONOSUPPORT
}

func _isPromiseOK(element *Promise) hestiaError.Error {
	return hestiaError.EPROTONOSUPPORT
}

func _isValueOK(element any) hestiaError.Error {
	return hestiaError.EPROTONOSUPPORT
}

func _setHTML(element *Object, html *[]byte) hestiaError.Error {
	return hestiaError.EPROTONOSUPPORT
}
