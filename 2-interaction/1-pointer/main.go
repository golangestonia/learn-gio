// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"

	"gioui.org/f32"        // f32 contains float32 points.
	"gioui.org/io/event"   // event contains general event information.
	"gioui.org/io/pointer" // pointer contains input/output for mouse and touch screens.
	"gioui.org/layout"     // layout is used for layouting widgets.
	"gioui.org/op"         // op is used for recording different operations.
	"gioui.org/op/paint"   // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp"   // qapp contains convenience funcs for this tutorial
	"github.com/golangestonia/learn-gio/qasset" // qasset contains convenience assets for this tutorial
)

var imageOp = paint.NewImageOp(qasset.Neutral)

var location = f32.Pt(300, 300)
var targetLocation = location

func main() {
	qapp.InputSize(func(ops *op.Ops, queue event.Queue, windowSize image.Point) {
		// register area for input events
		pointer.Rect(image.Rectangle{Max: windowSize}).Add(ops)

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
	})
}

func abs(s float32) float32 {
	if s < 0 {
		return -s
	}
	return s
}
