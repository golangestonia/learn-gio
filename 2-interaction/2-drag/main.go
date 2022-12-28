// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/f32"        // f32 contains float32 points.
	"gioui.org/gesture"    // gesture contains different gesture events
	"gioui.org/io/event"   // event contains general event information.
	"gioui.org/io/pointer" // pointer contains input/output for mouse and touch screens.
	"gioui.org/op"         // op is used for recording different operations.
	"gioui.org/op/clip"
	"gioui.org/op/paint" // paint contains operations for coloring.
	"gioui.org/unit"     // unit contains metric conversion

	"github.com/golangestonia/learn-gio/qapp"   // qapp contains convenience funcs for this tutorial
	"github.com/golangestonia/learn-gio/qasset" // qasset contains convenience assets for this tutorial
)

var imageOp = paint.NewImageOp(qasset.Neutral)

var location = f32.Pt(300, 300)
var drag gesture.Drag

func main() {
	qapp.Metric(func(ops *op.Ops, queue event.Queue, metric unit.Metric) {
		// handle drag events
		var dragOffset f32.Point
		for _, ev := range drag.Events(metric, queue, gesture.Both) {
			if ev.Type == pointer.Drag {
				dragOffset = ev.Position
			}
		}
		location = location.Add(dragOffset)

		// update the offset, must be after drag.Events
		defer op.Affine(f32.Affine2D{}.Offset(location)).Push(ops).Pop()

		// register image area for input events
		defer clip.Rect{Max: imageOp.Size()}.Push(ops).Pop()
		drag.Add(ops)

		// draw the image
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
