// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image/color"

	"gioui.org/f32"      // f32 contains float32 points.
	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

func main() {
	qapp.Render(func(ops *op.Ops) {
		// create a custom clip shape
		var p clip.Path
		p.Begin(ops)
		p.MoveTo(f32.Pt(30, 30))
		p.LineTo(f32.Pt(170, 170))
		p.LineTo(f32.Pt(80, 170))
		// the path must be closed
		p.Close()

		// set the clip to the outline
		defer clip.Outline{
			Path: p.End(),
		}.Op().Push(ops).Pop()

		// color the clip area:
		red := color.NRGBA{R: 0xFF, A: 0xFF}
		paint.ColorOp{Color: red}.Add(ops)
		paint.PaintOp{}.Add(ops)
	})
}
