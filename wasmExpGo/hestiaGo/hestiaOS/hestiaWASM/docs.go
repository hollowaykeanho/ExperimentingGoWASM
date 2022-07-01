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

// Package hestiaWASM is the abstracted elements interacting with DOM.
//
// The purpose of this package is to abstract and standardizes the DOM
// interactions with Go codes. Since `syscall/js` package is still experimental
// and it is subjected to change for future releases, it's best to keep the
// an adapter package here.
//
// PACKAGE USAGE
//
// Using this package directly is rare unless you plan to interact with
// Javascript directly. Otherwise, all hestiaGo packages are built on top of
// this hestiaASM adapter for maximum compatibility.
//
// WARNING: EXPERIMENTAL NATURE
//
// Due to the experimental nature of `syscall/js` package where is is the
// **ONLY** package to interact with DOM, there is **NO GUARENTEE** on API
// stability for this hestiaWASM.
//
// FOR WEBASSEMBLY BUILD ONLY
//
// As the package name implies, hestiaWASM is **ONLY meant for WASM build**.
// It only respects `target=wasm` or `CPU=wasm` build environment.
//
// While cross-platform compatibility is facilitated, all functions and objects
// are stubbed.
//
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
//   8. if output == ok { return hestiaError.OK }
//   9. not WASM build == unsupported { return hestiaError.EPFNOSUPPORT }
package hestiaWASM
