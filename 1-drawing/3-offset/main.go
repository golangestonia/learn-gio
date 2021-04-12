// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"

	"gioui.org/f32"      // f32 contains float32 points.
	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

func main() {
	qapp.Render(func(ops *op.Ops) {
		op.Offset(f32.Pt(100, 100)).Add(ops)
		clip.Rect{Max: image.Pt(100, 100)}.Add(ops)

		// color the clip area:
		red := color.NRGBA{R: 0xFF, A: 0xFF}
		paint.ColorOp{Color: red}.Add(ops)
		paint.PaintOp{}.Add(ops)
	})
}
