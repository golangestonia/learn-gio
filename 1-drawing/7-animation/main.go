// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"

	"gioui.org/f32"      // f32 contains float32 points.
	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp"   // qapp contains convenience funcs for this tutorial
	"github.com/golangestonia/learn-gio/qasset" // qasset contains convenience assets for this tutorial
)

var imageOp = paint.NewImageOp(qasset.Neutral)

var position float32

func main() {
	qapp.Render(func(ops *op.Ops) {
		position += 0.5
		op.Offset(f32.Pt(position, 0)).Add(ops)

		// the render needs to be called immediately again
		op.InvalidateOp{}.Add(ops)

		clip.Rect{Max: image.Pt(200, 200)}.Add(ops)
		imageOp.Add(ops)
		paint.PaintOp{}.Add(ops)
	})
}
