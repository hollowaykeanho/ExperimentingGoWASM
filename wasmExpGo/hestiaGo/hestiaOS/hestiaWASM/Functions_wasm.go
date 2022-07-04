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
	"unsafe"
)

const (
	id_JS_ADD_EVENT_LISTENER      = "addEventListener"
	id_JS_APPEND                  = "append"
	id_JS_CREATE_ELEMENT          = "createElement"
	id_JS_EVENT_BUBBLES           = "bubbles"
	id_JS_EVENT_CANCELABLE        = "cancelable"
	id_JS_EVENT_COMPOSED          = "composed"
	id_JS_EVENT_CURRENT_TARGET    = "currentTarget"
	id_JS_EVENT_DEFAULT_PREVENTED = "defaultPrevented"
	id_JS_EVENT_PHASE             = "eventPhase"
	id_JS_EVENT_PREVENT_DEFAULT   = "preventDefault"
	id_JS_EVENT_OPTION_CAPTURE    = "capture"
	id_JS_EVENT_OPTION_ONCE       = "once"
	id_JS_EVENT_OPTION_PASSIVE    = "passive"
	id_JS_EVENT_IS_TRUSTED        = "isTrusted"
	id_JS_EVENT_TARGET            = "target"
	id_JS_EVENT_TIMESTAMP         = "timeStamp"
	id_JS_EVENT_TYPE              = "type"
	id_JS_HTML                    = "innerHTML"
	id_JS_PROMISE                 = "Promise"
	id_JS_REMOVE_EVENT_LISTENER   = "removeEventListener"
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

func _addEventListener(element *Object, listener *EventListener) (err hestiaError.Error) {
	var options map[string]any
	var handler js.Func

	if IsObjectOK(element) != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	if IsEventListenerOK(listener) != hestiaError.OK {
		return hestiaError.ENOMEDIUM
	}

	// check if listener is already attached
	if listener.handler != nil {
		return hestiaError.EBADE
	}

	// create the Javascript compatible options list
	options = map[string]any{
		id_JS_EVENT_OPTION_CAPTURE: listener.Capture,
		id_JS_EVENT_OPTION_ONCE:    listener.Once,
		id_JS_EVENT_OPTION_PASSIVE: listener.Passive,
	}

	// create the Javascript compatible handler
	handler = js.FuncOf(func(this js.Value, args []js.Value) any {
		var obj js.Value

		// prevent default if set
		if listener.PreventDefault {
			args[0].Call(id_JS_EVENT_PREVENT_DEFAULT)
		}

		// convert the event parameters into Go format
		e := &Event{
			IsBubble:         args[0].Get(id_JS_EVENT_BUBBLES).Bool(),
			IsCancellable:    args[0].Get(id_JS_EVENT_CANCELABLE).Bool(),
			IsComposed:       args[0].Get(id_JS_EVENT_COMPOSED).Bool(),
			Phase:            EventPhase(args[0].Get(id_JS_EVENT_PHASE).Int()),
			DefaultPrevented: args[0].Get(id_JS_EVENT_DEFAULT_PREVENTED).Bool(),
			IsTrusted:        args[0].Get(id_JS_EVENT_IS_TRUSTED).Bool(),
			Timestamp:        args[0].Get(id_JS_EVENT_TIMESTAMP).Float(),
			Type:             args[0].Get(id_JS_EVENT_TYPE).String(),
		}

		if this.Equal(*(element.value)) {
			e.This = element
		} else {
			e.This = &Object{value: &this}
		}

		obj = args[0].Get(id_JS_EVENT_TARGET)
		if obj.Equal(*(element.value)) {
			e.Target = element
		} else {
			e.Target = &Object{value: &obj}
		}

		obj = args[0].Get(id_JS_EVENT_CURRENT_TARGET)
		if obj.Equal(*(element.value)) {
			e.CurrentTarget = element
		} else {
			e.CurrentTarget = &Object{value: &obj}
		}

		// execute the listener function
		go listener.Function(e)

		// return nothing since this is a JS function wrapper.
		return nil
	})

	// call the JS.addEventListener API
	element.value.Call(id_JS_ADD_EVENT_LISTENER,
		listener.Name,
		handler,
		options,
	)

	// save function for later release
	listener.handler = &Object{
		value:    element.value,
		function: &handler,
	}

	return hestiaError.OK
}

func _append(parent *Object, child *Object) hestiaError.Error {
	if IsObjectOK(parent) != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	if IsObjectOK(child) != hestiaError.OK {
		return hestiaError.ENOENT
	}

	parent.value.Call(id_JS_APPEND, *(child.value))

	return hestiaError.OK
}

func _createElement(name string) (child *Object, err hestiaError.Error) {
	if name == "" {
		return nil, hestiaError.ENODATA
	}

	parent := Document()

	ret := parent.value.Call(id_JS_CREATE_ELEMENT, name)

	return &Object{
		value: &ret,
	}, hestiaError.OK
}

func _execJSFunc(withRet bool, name string, args []any) (out any, err hestiaError.Error) {
	var ret js.Value
	var global *Object

	// validate all args are convertable to prevent possible panics
	for arg := range args {
		if IsTypeConvertable(arg) != hestiaError.OK {
			return nil, hestiaError.EINVAL
		}
	}

	// get Javascript function
	global = Global()
	ret = global.value.Get(name)
	if ret.Type() != js.TypeFunction {
		return nil, hestiaError.EPROTOTYPE
	}

	// Invoke Javascript function
	ret = ret.Invoke(args...)
	err = hestiaError.OK
	if !withRet {
		out = nil
		goto done
	}

	// Convert return value to compatible Go format
	switch ret.Type() {
	case js.TypeBoolean:
		out = ret.Bool()
	case js.TypeNumber:
		out = ret.Float()
	case js.TypeNull:
		out = nil
	case js.TypeUndefined:
		out = nil
	case js.TypeObject:
		out = "<Javascript Object>"
	case js.TypeFunction:
		out = "<Javascript Function>"
	case js.TypeString:
		fallthrough
	default:
		out = ret.String()
	}

done:
	return out, err
}

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

func _removeEventListener(element *Object, listener *EventListener) hestiaError.Error {
	var options map[string]any

	if IsObjectOK(element) != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	if IsEventListenerOK(listener) != hestiaError.OK {
		return hestiaError.ENOMEDIUM
	}

	// check if listener is already free
	if listener.handler == nil {
		return hestiaError.EBADE
	}

	if listener.handler.function == nil {
		return hestiaError.EBADE
	}

	// create the Javascript compatible options list
	options = map[string]any{
		id_JS_EVENT_OPTION_CAPTURE: listener.Capture,
	}

	// call the JS.removeEventListener API
	element.value.Call(id_JS_REMOVE_EVENT_LISTENER,
		listener.Name,
		*(listener.handler.function),
		options,
	)

	// release save function
	listener.handler = nil

	return hestiaError.OK
}

func _setHTML(element *Object, html *[]byte) hestiaError.Error {
	if html == nil {
		return hestiaError.ENODATA
	}

	if IsObjectOK(element) != hestiaError.OK {
		return hestiaError.EOWNERDEAD
	}

	element.value.Set(id_JS_HTML, string(*html))

	return hestiaError.OK
}

// NOTE: all functions below are sub-functions. Please use the global version
// since it has proper guarding like `nil` object checking.

func _goPromise(promise *Promise) (err hestiaError.Error) {
	var jsFunc js.Func

	// obtain promise object from global
	if promise.object == nil {
		promise.object = Get(Global(), id_JS_PROMISE)
	}

	// generate promise function
	jsFunc = js.FuncOf(func(this js.Value, args []js.Value) any {
		handler := __newGenericJSPromiseHandler(promise)
		return promise.object.value.New(*handler)
	})
	promise.object.function = &jsFunc

	// expose Promise function to Javascript
	Global().value.Set(promise.Name, *(promise.object.function))

	return hestiaError.OK
}

func __newGenericJSPromiseHandler(promise *Promise) *js.Func {
	var handler js.Func

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

			if ret == hestiaError.OK {
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

	return &handler
}

func _isEventListenerOK(element *EventListener) hestiaError.Error {
	if element.Name == "" {
		return hestiaError.EBADF
	}

	if element.Function == nil {
		return hestiaError.ENOENT
	}

	return hestiaError.OK
}

func _isObjectOK(element *Object) hestiaError.Error {
	if element.value == nil {
		return hestiaError.ENOENT
	}

	return hestiaError.OK
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

func _isTypeConvertable(element any) hestiaError.Error {
	switch element.(type) {
	case js.Value:
	case js.Func:
	case nil:
	case bool:
	case int, int8, int16, int32, int64:
	case uint, uint8, uint16, uint32, uint64:
	case uintptr:
	case unsafe.Pointer:
	case float32, float64:
	case string:
	case []any:
	case map[string]any:
	default:
		return hestiaError.EPROTOTYPE
	}

	return hestiaError.OK
}
