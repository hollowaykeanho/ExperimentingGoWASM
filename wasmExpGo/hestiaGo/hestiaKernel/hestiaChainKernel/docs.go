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

// hestiaChainKernel is the chain-executing function blocks kernel structure.
//
// This kernal approach orignated from chain design architecture where function
// execution blocks are executed and next block is determined on-the-fly. In
// other word, the previous block decides and setup the next function block
// right before it returns its input value. This is suitable for chaotic
// event-driven oriented executions such as but not limited to user-driven
// activities or interrupt-driven activities.
//
// Chain is designed to be thread-safe in order to work both stand-alone as a
// kernel in non-OS environment or on top of another kernel structure
// (e.g. App).
//
// This kernel structure relies heavily on `func(any) any` function pattern
// where the previous block's output value is the input of the next block. It
// **ONLY HANDLES ONE (1)** parameter so if you need multiple ones, consider
// defining a data `struct` or using a `map` data list. To avoid restricting
// the kernel to a certain data type, the chain shall accept `any` data type and
// it's up to the user to assert the data type inside the function block. One
// easy way is:
//       func myBlock(arg any) any {
//           out, ok := arg.(MyType)
//           if !ok {
//               return nil
//           }
//
//           ...
//       }
package hestiaChainKernel
