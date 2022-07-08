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

package hestiaUI

type CSSVariable struct {
	Key   string
	Value string
}

type CSSVarList []*CSSVariable

// CSS Variables are the naming constant for Key in CSSVariable data structure.
//
// All variable names are listed here to ensure all specific UI packages can
// import without getting into circular dependencies with each others.
const (
	// hestiaCoreUI - standard constants
	CSS_VAR_Z_MAX = "--z-index-max"

	// hestiaCoreUI - <html>
	CSS_VAR_HTML_BORDER_BOX = "--html-border-box"

	// hestiaCoreUI - <body>
	CSS_VAR_BODY_GAP           = "--body-gap"
	CSS_VAR_BODY_GRID          = "--body-grid"
	CSS_VAR_BODY_GRID_VERTICAL = "--body-grid-vertical"

	// hestiaCoreUI - <main>
	CSS_VAR_MAIN_Z_INDEX = "--main-z-index"
	CSS_VAR_MAIN_PADDING = "--main-padding"
)

// CSS Values are the constnat values used in CSS across all UI packages.
const (
	CSS_VALUE_Z_INDEX_MAX = "2147483647"

	CSS_VALUE_LAYOUT_SEGMENT_TOP     = "topSegment"
	CSS_VALUE_LAYOUT_SEGMENT_LEFT    = "leftSegment"
	CSS_VALUE_LAYOUT_SEGMENT_RIGHT   = "rightSegment"
	CSS_VALUE_LAYOUT_SEGMENT_BOTTOM  = "bottomSegment"
	CSS_VALUE_LAYOUT_SEGMENT_CONTENT = "contentSegment"
)
