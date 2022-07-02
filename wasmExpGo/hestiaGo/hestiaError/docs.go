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
// list to ensure a memory-size constrained environment for TinyGo can operate
// properly without being overburden-ed by Go's string error approach.
//
//
// WHY LINUX ERROR CODES?
//
// Since most server systems in the world are using Linux kernel system and
// TinyGo is using Linux kernel and it deployed operating system for many of
// their respective applications, it is the best reference point. Besides,
// microcontrollers implementation is a baremetal (non-OS) approach so the
// kernel-level insight is best suited.
//
// Some mild wording were smartly applied to the text representations of the
// error code for flexible reusability in different aspects and domains instead
// to the original one. When used, you should prepare your own intepretation
// with association as close to this list as possible.
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
// numbers away from this list. HestiaError recommends `uint16` to consider
// your custom codes (beyond this list).
//
// `uint8` is not recommended due to the consideration of `int8`, where the
// maximum number is 124. The list already exceeded that capacity.
//
// `int` or `uint` is **strongly not recommended** since HestiaGo supports
// various bit-size CPU architecture (e.g. 8-bit microcontroller to 64-bit
// x86 CPU).
//
//
// WHY ALL CONSTANTS?
//
// Mainly to allow Go and TinyGo compilers to shed unused ones to reduce your
// compiled binary size.
//
// This compiler "code shedding" ability is also the reason why the
// `Interpret(...)` function idea was seriously considered but ultimately
// disposed of due to the function itself is fat and you're likely going to have
// your interpretation function anyway.
//
//
// WHY NOT Go.Error TYPE?
//
// Same reason with TinyGo operations. However, you can always create the Go
// error object in Go using the associated string constant (by prefixing
// `ERROR_`) or your own message based on the error code.
package hestiaError
