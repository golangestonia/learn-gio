// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app" // app contains Window handling.
	// gesture contains different gesture event handling.
	// pointer contains input/output for mouse and touch screens.
	"gioui.org/io/system" // system is used for system events (e.g. closing the window).
	"gioui.org/layout"    // layout is used for layouting widgets.
	"gioui.org/op"        // op is used for recording different operations.
	"gioui.org/op/clip"   // clip contains operations for clipping painting area.
	"gioui.org/op/paint"  // paint contains operations for coloring.
)

func render(gtx layout.Context) {
	// Center
	// N E S W
	// NE SE SW NW
	layout.NE.Layout(gtx,
		Box{
			Size:  image.Pt(150, 150),
			Color: color.NRGBA{R: 0xFF, A: 0xFF},
		}.Layout,
	)
}

type Box struct {
	Size  image.Point
	Color color.NRGBA
}

func (box Box) Layout(gtx layout.Context) layout.Dimensions {
	defer op.Save(gtx.Ops).Load()

	size := gtx.Constraints.Constrain(box.Size)

	clip.Rect{Max: size}.Add(gtx.Ops)
	paint.ColorOp{Color: box.Color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

	return layout.Dimensions{
		Size: size,
	}
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
			render(gtx)
			// render and handle the operations from the UI.
			e.Frame(gtx.Ops)

		// this is sent when the application is closed.
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}
