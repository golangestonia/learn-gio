// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/f32"      // f32 contains float32 points.
	"gioui.org/io/event" // event contains general event information.
	"gioui.org/io/key"   // key contains input/output for keyboards.
	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp"   // qapp contains convenience funcs for this tutorial
	"github.com/golangestonia/learn-gio/qasset" // qasset contains convenience assets for this tutorial
)

var imageOp = paint.NewImageOp(qasset.Neutral)

const speed = 50

var location = f32.Pt(300, 300)

func main() {
	qapp.Input(func(ops *op.Ops, queue event.Queue) {
		// keep the focus, since only one thing can
		key.FocusOp{Tag: &location}.Add(ops)
		// register tag &location as reading input
		key.InputOp{Tag: &location}.Add(ops)

		// read events from input event queue
		for _, ev := range queue.Events(&location) {
			// figure out, which event it was
			switch ev := ev.(type) {
			case key.Event:
				if ev.State == key.Press {
					switch ev.Name {
					case key.NameLeftArrow:
						location.X -= speed
					case key.NameUpArrow:
						location.Y -= speed
					case key.NameRightArrow:
						location.X += speed
					case key.NameDownArrow:
						location.Y += speed
					}
				}
			}
		}

		op.Offset(location).Add(ops)
		imageOp.Add(ops)
		paint.PaintOp{}.Add(ops)
	})
}
