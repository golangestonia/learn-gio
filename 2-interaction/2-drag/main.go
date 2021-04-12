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
	"gioui.org/gesture"    // gesture contains different gesture events
	"gioui.org/io/event"   // event contains general event information.
	"gioui.org/io/pointer" // pointer contains input/output for mouse and touch screens.
	"gioui.org/io/system"  // system is used for system events (e.g. closing the window).
	"gioui.org/layout"     // layout is used for layouting widgets.
	"gioui.org/op"         // op is used for recording different operations.
	"gioui.org/op/paint"   // paint contains operations for coloring.
	"gioui.org/unit"       // unit contains metric conversion
)

var location = f32.Pt(300, 300)
var drag gesture.Drag

func render(ops *op.Ops, queue event.Queue, size image.Point, metric unit.Metric) {
	// handle drag events
	var dragOffset f32.Point
	for _, ev := range drag.Events(metric, queue, gesture.Both) {
		if ev.Type == pointer.Drag {
			dragOffset = ev.Position
		}
	}
	location = location.Add(dragOffset)

	// update the offset, must be after drag.Events
	op.Offset(location).Add(ops)

	// register image area for input events
	pointer.Rect(image.Rectangle{Max: imageOp.Size()}).Add(ops)
	drag.Add(ops)

	// draw the image
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
			render(gtx.Ops, gtx.Queue, gtx.Constraints.Max, gtx.Metric)
			// render and handle the operations from the UI.
			e.Frame(gtx.Ops)

		// this is sent when the application is closed.
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}
