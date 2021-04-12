// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image/color"

	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

func main() {
	qapp.Render(func(ops *op.Ops) {
		red := color.NRGBA{R: 0xFF, A: 0xFF}
		// ColorOp sets the brush for painting.
		paint.ColorOp{Color: red}.Add(ops)
		// PaintOp paints the configured.
		paint.PaintOp{}.Add(ops)
	})
}
