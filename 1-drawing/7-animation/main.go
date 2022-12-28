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
		// Note, if we use op.Offset, the image will be pixel-snapped.
		// The behavior we want, depends on the context.
		defer op.Affine(f32.Affine2D{}.Offset(f32.Pt(position, 0))).Push(ops).Pop()

		// the render needs to be called immediately again
		op.InvalidateOp{}.Add(ops)

		defer clip.Rect{Max: image.Pt(200, 200)}.Push(ops).Pop()
		imageOp.Add(ops)
		paint.PaintOp{}.Add(ops)
	})
}
