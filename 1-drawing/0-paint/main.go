// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"       // app contains Window handling.
	"gioui.org/io/system" // system is used for system events (e.g. closing the window).
	"gioui.org/layout"    // layout is used for layouting widgets.
	"gioui.org/op"        // op is used for recording different operations.
	"gioui.org/op/paint"  // paint contains operations for coloring.
)

func render(ops *op.Ops, size image.Point) {
	red := color.NRGBA{R: 0xFF, A: 0xFF}
	// ColorOp sets the brush for painting.
	paint.ColorOp{Color: red}.Add(ops)
	// PaintOp paints the configured.
	paint.PaintOp{}.Add(ops)
}

/* usual setup code */

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
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
			render(gtx.Ops, gtx.Constraints.Max)
			// render and handle the operations from the UI.
			e.Frame(gtx.Ops)

		// this is sent when the application is closed.
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}
