// Copyright 2022 "Holloway" Chew, Kean Ho <hollowaykeanho@gmail.com>
// Copyright 2021 Chan Wen Xu <chan@wenxu.dev>
// Copyright 2020 Alessandro Segala <twitter: @ItalyPaleAle>
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

package hestiaWASM

import (
	"hestiaGo/hestiaError"
	"syscall/js"
)

const (
	id_JS_PROMISE = "Promise"

	hestiaERROR_OK = 0
)

// RETURN ERROR CODES
//
// HestiaWASM tries to standardizes its return error codes based on syscall/js
// and Javascript nature. That way, it can stabilizes all upper layers'
// developed codes and preparing itself for operations. Here are the list of
// known error codes:
//
//   1. if input parent == nil { return hestiaError.EOWNERDEAD }
//   2. if input child == nil { return hestiaError.ENOENT }
//   3. if input (any) arguments == faulty { return hestiaError.ENOENT }
//   4. if input (any) arguments == missing { return hestiaError.ENODATA }
//   5. if output == missing { return hestiaError.ENOPROTOOPT }
//   6. if output == bad { return hestiaError.EPROTO }
//   7. if output == unsupported { return hestiaError.EPROTONOSUPPORT }
//   7. if output == ok { return hestiaError.OK }

func _get(parent *Object, query string) *Object {
	if query == "" {
		return nil
	}

	if IsObjectOK(parent) != hestiaError.OK {
		return nil
	}

	ret := parent.value.Get(query)

	return &Object{
		value: &ret,
	}
}

func _createElement(name string) (child *Object, err hestiaError.Error) {
	if name == "" {
		return nil, hestiaError.ENODATA
	}

	parent := Document()

	ret := parent.value.Call("createElement", name)

	return &Object{
		value: &ret,
	}, hestiaError.OK
}

func _appendChild(parent *Object, child *Object) hestiaError.Error {
	if IsObjectOK(parent) != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	if IsObjectOK(child) != hestiaError.OK {
		return hestiaError.ENOENT
	}

	parent.value.Call("appendChild", *(child.value))

	return hestiaError.OK
}

func _setHTML(element *Object, html *[]byte) hestiaError.Error {
	if html == nil {
		return hestiaError.ENODATA
	}

	if IsObjectOK(element) != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	element.value.Set("innerHTML", string(*html))

	return hestiaError.OK
}

// NOTE: all functions below are sub-functions. Please use the global version
// since it has proper guarding like `nil` object checking.

func _goPromise(promise *Promise) (err hestiaError.Error) {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := __newJSPromise(promise)

		promise.object = Get(Global(), id_JS_PROMISE)
		return promise.object.value.New(handler)
	})

	Global().value.Set(promise.Name, jsFunc)

	return hestiaError.OK
}

func __newJSPromise(promise *Promise) (handler js.Func) {
	handler = js.FuncOf(func(this js.Value, args []js.Value) any {
		switch {
		case len(args) < 2:
			promise.Reject(hestiaError.ENOTRECOVERABLE)
			return nil
		case args[0].Type() != js.TypeFunction,
			args[1].Type() != js.TypeFunction:
			promise.Reject(hestiaError.ENOTRECOVERABLE)
			return nil
		default:
		}

		go func() {
			ret := promise.Func()

			if ret == hestiaERROR_OK {
				// resolve
				args[0].Invoke(js.ValueOf(promise.Resolve()))
			} else {
				//reject
				args[1].Invoke(js.ValueOf(promise.Reject(hestiaError.OK)))
			}

			handler.Release()
		}()

		return nil
	})

	return handler
}

func _isPromiseOK(element *Promise) hestiaError.Error {
	if element.Name == "" {
		return hestiaError.EBADF
	}

	if element.Func == nil {
		return hestiaError.ENOENT
	}

	if element.Resolve == nil {
		return hestiaError.ENOPROTOOPT
	}

	if element.Reject == nil {
		return hestiaError.ENOMEDIUM
	}

	return hestiaError.OK
}

func _isObjectOK(element *Object) hestiaError.Error {
	if element.value == nil {
		return hestiaError.ENOENT
	}

	return hestiaError.OK
}
