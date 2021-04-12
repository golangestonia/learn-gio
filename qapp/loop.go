// SPDX-License-Identifier: Unlicense OR MIT

package qapp

import (
	"log"
	"os"

	"gioui.org/app"       // app contains Window handling.
	"gioui.org/io/event"  // event contains general event information.
	"gioui.org/io/system" // system is used for system events (e.g. closing the window).
	"gioui.org/layout"    // layout is used for layouting widgets.
	"gioui.org/op"        // op is used for recording different operations.
)

// Render is a utility to start a rendering gio app.
func Render(fn func(op *op.Ops)) {
	go func() {
		w := app.NewWindow()
		// ops will be used to encode different operations.
		var ops op.Ops

		// listen for events happening on the window.
		for e := range w.Events() {
			// detect the type of the event.
			switch e := e.(type) {
			// this is sent when the application should re-render.
			case system.FrameEvent:
				// gtx is used to pass around rendering and event information.
				gtx := layout.NewContext(&ops, e)
				// render content
				fn(gtx.Ops)
				// render and handle the operations from the UI.
				e.Frame(gtx.Ops)

			// this is sent when the application is closed.
			case system.DestroyEvent:
				if e.Err != nil {
					log.Println(e.Err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}()
	app.Main()
}

// Input is a utility to start a rendering and input gio app.
func Input(fn func(op *op.Ops, q event.Queue)) {
	go func() {
		w := app.NewWindow()
		// ops will be used to encode different operations.
		var ops op.Ops

		// listen for events happening on the window.
		for e := range w.Events() {
			// detect the type of the event.
			switch e := e.(type) {
			// this is sent when the application should re-render.
			case system.FrameEvent:
				// gtx is used to pass around rendering and event information.
				gtx := layout.NewContext(&ops, e)
				// render content
				fn(gtx.Ops, gtx.Queue)
				// render and handle the operations from the UI.
				e.Frame(gtx.Ops)

			// this is sent when the application is closed.
			case system.DestroyEvent:
				if e.Err != nil {
					log.Println(e.Err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}()
	app.Main()
}

// Layout is a utility to start a layouting gio app.
func Layout(lay func(gtx layout.Context) layout.Dimensions) {
	go func() {
		w := app.NewWindow()
		// ops will be used to encode different operations.
		var ops op.Ops

		// listen for events happening on the window.
		for e := range w.Events() {
			// detect the type of the event.
			switch e := e.(type) {
			// this is sent when the application should re-render.
			case system.FrameEvent:
				// gtx is used to pass around rendering and event information.
				gtx := layout.NewContext(&ops, e)
				// render content
				lay(gtx)
				// render and handle the operations from the UI.
				e.Frame(gtx.Ops)

			// this is sent when the application is closed.
			case system.DestroyEvent:
				if e.Err != nil {
					log.Println(e.Err)
					os.Exit(1)
				}
				os.Exit(0)
			}
		}
	}()
	app.Main()
}
