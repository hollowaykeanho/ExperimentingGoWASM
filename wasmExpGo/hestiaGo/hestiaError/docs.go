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

// Package hestiaError is a compact memory-size sensitive error list.
//
// The sole purpose for package is to enable Go and TinyGo to operate in-sync
// when communicating an error. This is done by re-visiting Linux error numbers
// list to ensure memory-size constrained environment for TinyGo can operate
// properly without being overburden-ed by Go's string error approach.
//
//
// WHY LINUX ERROR CODE?
//
// Since most server systems in the world are using Linux kernel system and
// TinyGo is also used to build kernel and operating system for
// microcontrollers, Linux kernel system thus becomes the best reference for
// this type of requirement.
//
// Some mild wording were smartly applied to the text intepretations for
// flexible reusability in different aspects and domains instead of a fixed
// one.
//
//
// NOT ALL ERROR CODES ARE AVAILABLE
//
// Please keep in mind that not all the error codes are ported into this
// package. Some error codes are senseless as they're specific to Linux kernel
// usage.
//
// However, to ensure maximum interpolatibility, all error numbers retained
// their former number.
//
//
// POSITIVE OR NEGATIVE NUMBER
//
// Entirely your choice. Using positive number allows you to expand greater
// numbers away from this list.
//
//
// WHY ALL CONSTANTS?
//
// Mainly to allow Go and TinyGo compilers to shed unused codes.
//
//
// WHY NOT Go.Error TYPE?
//
// Same reason with TinyGo operations. You can always create the error object
// in Go using the associated string constant (by prefixing `ERROR_` ) .
package hestiaError
