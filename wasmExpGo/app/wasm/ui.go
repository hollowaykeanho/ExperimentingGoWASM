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

package main

import (
	"fmt"

	"hestiaGo/hestiaKernel/hestiaChainKernel"
	"hestiaGo/hestiaOS"
	"hestiaGo/hestiaOS/hestiaWASM"
	"hestiaGo/hestiaUI/hestiaCoreUI"
)

type ui struct {
	event    *hestiaWASM.Event
	kernel   *hestiaChainKernel.Kernel
	button   *hestiaWASM.Object
	listener *hestiaWASM.EventListener
}

func uiInit() {
	var controller *ui

	// create the UI controller
	controller = &ui{
		kernel: &hestiaChainKernel.Kernel{},
		listener: &hestiaWASM.EventListener{
			Name: "click",
			Function: func(e *hestiaWASM.Event) {
				controller.event = e
				_ = hestiaChainKernel.Trigger(controller.kernel)
			},
			PreventDefault: true,
		},
	}

	// render base UI for first interaction
	controller.button, _ = hestiaWASM.CreateElement("button")
	html := []byte("Render WASM Contents")
	_ = hestiaWASM.SetHTML(controller.button, &html)
	_ = hestiaWASM.AddEventListener(controller.button, controller.listener)
	_ = hestiaWASM.Append(hestiaWASM.Body(), controller.button)

	// generate and debug CSS
	css := hestiaCoreUI.CSS(hestiaCoreUI.CSSVariables(), false)
	fmt.Printf("Core CSS BELOW:\n%s[END]", css)

	// start chain server
	hestiaChainKernel.Start(controller.kernel, func(arg any) (out any) {
		signal, ok := arg.(uint16)
		if !ok {
			return nil
		}

		switch signal {
		case hestiaOS.SIGNAL_SIGCONT:
			hestiaChainKernel.SetNext(controller.kernel, _sourceUIChanges)
		case hestiaOS.SIGNAL_SIGSTOP,
			hestiaOS.SIGNAL_SIGINT,
			hestiaOS.SIGNAL_SIGTERM,
			hestiaOS.SIGNAL_SIGKILL:
		}

		return controller
	}, 5)
}

func _sourceUIChanges(arg any) any {
	controller := __convertArgument(arg)

	// execute function
	fmt.Printf("analyzing and sourcing changes from rendered UI!\n")
	fmt.Printf("Got event data: %#v\n", controller.event)

	// chain next event
	hestiaChainKernel.SetNext(controller.kernel, _renderUIChanges)

	// end function block
	return controller
}

func _renderUIChanges(arg any) any {
	controller := __convertArgument(arg)

	// execute function
	body := hestiaWASM.Body()
	tag, _ := hestiaWASM.CreateElement("h2")
	html := []byte("button content rendered here!")
	_ = hestiaWASM.SetHTML(tag, &html)
	_ = hestiaWASM.Append(body, tag)

	// chain next event
	hestiaChainKernel.SetNext(controller.kernel, _removeUIFunction)

	// end function block
	return controller
}

func _removeUIFunction(arg any) any {
	controller := __convertArgument(arg)

	// execute function
	fmt.Printf("removing the listener 1st-time\n")
	_ = hestiaWASM.RemoveEventListener(controller.button, controller.listener)

	fmt.Printf("removing the listener 2nd-time\n")
	_ = hestiaWASM.RemoveEventListener(controller.button, controller.listener)

	fmt.Printf("removing the listener 3rd-time\n")
	_ = hestiaWASM.RemoveEventListener(controller.button, controller.listener)

	fmt.Printf("removing the listener 4th-time\n")
	_ = hestiaWASM.RemoveEventListener(controller.button, controller.listener)

	// chain next event
	// DONE - no more chaining since UI is dead. Stopping controller as
	//        well.
	_ = hestiaChainKernel.Stop(controller.kernel)

	// end function block
	return controller
}

func __convertArgument(arg any) (controller *ui) {
	var ok bool

	if controller, ok = arg.(*ui); !ok {
		return nil
	}

	return controller
}
