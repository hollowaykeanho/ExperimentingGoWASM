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

// Get obtains a child element from a given parent element.
//
// It accepts the following parameters:
//   1. `parent` - the parent hestiaWASM.Object (e.g. highest is Global)
//   2. `query` - the name of your request.
//
// It shall returns:
//   1. `nil` - any parameters that are missing or empty
//   2. hestiaWASM.Object - a queried child element object including its `null`
//                          or `nil` Javascript nature.
func Get(parent *Object, query string) *Object {
	return _get(parent, query)
}

// CreateElement creates a new Javascript element from Document object.
//
// It accepts the following parameters:
//   1. `name` - name of the element
//
// It shall returns:
//   1. `nil`, hestiaError.Error - any parameters that are missing or empty
//   2. hestiaWASM.Object, `hestiaError.OK` - a created element
func CreateElement(name string) (*Object, hestiaError.Error) {
	return _createElement(name)
}

// AppendChild appends a child element to a parent element.
//
// It accepts the following parameters:
//   1. `parent` - the element to receive new element.
//   2. `child` - the element for appending.
//
// It shall returns:
//   1. hestiaError.Error - any parameters that are missing or empty
//   2. hestiaError.OK - operation successful.
func AppendChild(parent *Object, child *Object) hestiaError.Error {
	return _appendChild(parent, child)
}

// GoPromise registers a given Promise into Javascript function.
//
// This function only registers the given Promise into Javascript domain making
// it readily to be called and executed.
//
// Its return value here is meant to report the registration status only.
//
// It accepts the following parameters:
//   1. `promise` - the hestiaWASM.Promise to execute.
//
// It shall returns:
//   1. hestiaError.OK | `0` - scheduling was successful.
//   2. All hestiaErrors from `IsPromiseOK()` - failed usability test.
func GoPromise(promise *Promise) (err hestiaError.Error) {
	err = IsPromiseOK(promise)
	if err != hestiaError.OK {
		return err
	}

	return _goPromise(promise)
}

// SetHTML applies a given HTML codes into a given element's InnerHTML.
//
// It accepts the following parameters:
//   1. `element` - the element to receive the HTML codes.
//   2. `html` - the pointer of the byte slice containing the HTML codes.
//
// It shall returns:
//   1. hestiaError.Error - any parameters that are missing or empty
//   2. hestiaError.OK - operation successful.
func SetHTML(element *Object, html *[]byte) hestiaError.Error {
	return _setHTML(element, html)
}

// IsPromiseOK checks a hestiaWASM.Promise is a stub or is operable.
//
// It accepts the following parameters:
//   1. `element` - the hestiaWASM.Promise to inspect
//
// It shall returns:
//   1. hestiaError.OK | `0` - The Promise object is operable.
//   2. hestiaError.EOWNERDEAD - The given Promise object is `nil`.
//   3. hestiaError.ENOENT - The Promise.Func property is `nil`.
//   4. hestiaError.ENOPROTOOPT - The Promise.Resolve property is `nil`.
//   5. hestiaError.ENOMEDIUM - The Promise.Reject property is `nil`.
//   6. hestiaError.EBADF - The Promise.Name property is empty (`""`).
//   7. hestiaError.EPROTONOSUPPORT - operating in a non-WASM CPU.
func IsPromiseOK(element *Promise) hestiaError.Error {
	if element == nil {
		return hestiaError.EOWNERDEAD
	}

	return _isPromiseOK(element)
}

// IsObjectOK checks a hestiaWASM.Object is a stub or is operable.
//
// It accepts the following parameters:
//   1. `element` - the hestiaWASM.Object to inspect
//
// It shall returns:
//   1. hestiaError.EOWNERDEAD - given `element` object is `nil`.
//   2. hestiaError.OK | `0` - [GOOD] the object is operable.
//   3. hestiaError.Error - [BAD] the object is not operable.
func IsObjectOK(element *Object) hestiaError.Error {
	if element == nil {
		return hestiaError.EOWNERDEAD
	}

	return _isObjectOK(element)
}