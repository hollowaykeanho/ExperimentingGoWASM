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

package hestiaCoreUI

import (
	"hestiaGo/hestiaUI"
)

func CSSVariables() *hestiaUI.CSSVarList {
	return &hestiaUI.CSSVarList{
		&hestiaUI.CSSVariable{
			Key:   hestiaUI.CSS_VAR_Z_MAX,
			Value: hestiaUI.CSS_VALUE_Z_INDEX_MAX,
		},
		&hestiaUI.CSSVariable{
			Key:   hestiaUI.CSS_VAR_HTML_BORDER_BOX,
			Value: "border-box",
		},
		&hestiaUI.CSSVariable{
			Key:   hestiaUI.CSS_VAR_BODY_GAP,
			Value: "0",
		},
		&hestiaUI.CSSVariable{
			Key: hestiaUI.CSS_VAR_BODY_GRID,
			Value: hestiaUI.CSS_VALUE_LAYOUT_SEGMENT_CONTENT + ` auto
	` + hestiaUI.CSS_VALUE_LAYOUT_SEGMENT_TOP + ` minmax(0 , auto)
	` + hestiaUI.CSS_VALUE_LAYOUT_SEGMENT_LEFT + ` minmax(0, auto)
	` + hestiaUI.CSS_VALUE_LAYOUT_SEGMENT_RIGHT + ` minmax(0, auto)
	` + hestiaUI.CSS_VALUE_LAYOUT_SEGMENT_BOTTOM + ` minmax(0, auto)
	/ 100%`,
		},
		&hestiaUI.CSSVariable{
			Key:   hestiaUI.CSS_VAR_MAIN_Z_INDEX,
			Value: "6",
		},
		&hestiaUI.CSSVariable{
			Key:   hestiaUI.CSS_VAR_MAIN_PADDING,
			Value: "3.5rem",
		},
	}
}

func CSS(variables *hestiaUI.CSSVarList, onlyVariables bool) (out string) {
	var i int
	var v *hestiaUI.CSSVariable

	// prepend variables if requested
	if variables != nil {
		for i, v = range *variables {
			if i != 0 {
				out += "\n"
			}

			out += v.Key + ": " + v.Value + ";"
		}
	}

	if onlyVariables {
		return out
	}

	// render component's CSS
	out += `
* {
	width: 100%;
	max-width: 100%;

	margin: 0 auto;
	padding: 0;
	vertical-align: middle;

	text-align: center;

	animation: .8s linear 0s infinite normal;
}

html {
	font-size: 62.5%; // 1.6rem = 16px

	height: 100%;
	height: calc(100vh - calc(100vh - 100%));

	box-sizing: var(` + hestiaUI.CSS_VAR_HTML_BORDER_BOX + `);
}

html,
body {
	margin: 0;
	padding: 0;
}

body {
	min-height: 100vh;
	display: grid;
	gap: var(` + hestiaUI.CSS_VAR_BODY_GAP + `);
	grid: var(` + hestiaUI.CSS_VAR_BODY_GRID + `);
}

main {
	z-index: calc(` + hestiaUI.CSS_VALUE_Z_INDEX_MAX + ` -
		var(` + hestiaUI.CSS_VAR_MAIN_Z_INDEX + `);
	padding: var(` + hestiaUI.CSS_VAR_MAIN_PADDING + `);
	grid-area: ` + hestiaUI.CSS_VALUE_LAYOUT_SEGMENT_CONTENT + `;
}
`

	// return final output
	return out
}
