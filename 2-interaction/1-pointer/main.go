// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"
	"os"

	"gioui.org/app"        // app contains Window handling.
	"gioui.org/f32"        // f32 contains float32 points.
	"gioui.org/io/event"   // event contains general event information.
	"gioui.org/io/pointer" // pointer contains input/output for mouse and touch screens.
	"gioui.org/io/system"  // system is used for system events (e.g. closing the window).
	"gioui.org/layout"     // layout is used for layouting widgets.
	"gioui.org/op"         // op is used for recording different operations.
	"gioui.org/op/paint"   // paint contains operations for coloring.
)

var location = f32.Pt(300, 300)
var targetLocation = location

func render(ops *op.Ops, queue event.Queue, size image.Point) {
	// register area for input events
	pointer.Rect(image.Rectangle{Max: size}).Add(ops)

	// register the area for pointer events
	pointer.InputOp{
		Tag:   &location,
		Types: pointer.Press,
	}.Add(ops)

	// read events from input event queue
	for _, ev := range queue.Events(&location) {
		// figure out, which event it was
		switch ev := ev.(type) {
		case pointer.Event:
			if ev.Type == pointer.Press {
				targetLocation = ev.Position
			}
		}
	}

	// move slightly towards the target location
	if delta := targetLocation.Sub(location); abs(delta.X) > 1 || abs(delta.Y) > 1 {
		delta.X *= 0.1
		delta.Y *= 0.1
		location = location.Add(delta)

		// ensure we animate this
		op.InvalidateOp{}.Add(ops)
	}

	op.Offset(location).Add(ops)

	imageSize := imageOp.Size().Div(-2)
	op.Offset(layout.FPt(imageSize)).Add(ops)

	imageOp.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func abs(s float32) float32 {
	if s < 0 {
		return -s
	}
	return s
}

/* usual setup code */

//go:embed neutral.png
var imageData []byte

var imageOp = func() paint.ImageOp {
	m, err := png.Decode(bytes.NewReader(imageData))
	if err != nil {
		panic(err)
	}
	return paint.NewImageOp(m)
}()

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
			render(gtx.Ops, gtx.Queue, gtx.Constraints.Max)
			// render and handle the operations from the UI.
			e.Frame(gtx.Ops)

		// this is sent when the application is closed.
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}
