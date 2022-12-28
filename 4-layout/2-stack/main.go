// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/font/gofont" // gofont is used for loading the default font.
	"gioui.org/layout"      // layout is used for layouting widgets.
	"gioui.org/op/clip"     // clip contains operations for clipping painting area.
	"gioui.org/op/paint"    // paint contains operations for coloring.
	"gioui.org/unit"
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var Theme = material.NewTheme(gofont.Collection())

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		return layout.Stack{
			Alignment: layout.Center,
		}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				size := gtx.Constraints.Max
				smallest := min(size.X, size.Y)

				size.X, size.Y = smallest, smallest
				shape := clip.Ellipse{Max: size}
				paint.FillShape(gtx.Ops, Theme.ContrastBg, shape.Op(gtx.Ops))

				return layout.Dimensions{
					Size: size,
				}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				size := gtx.Constraints.Max
				textSize := min(size.X, size.Y) * 3 / 4
				if textSize > 1023 {
					// There's a limit how large we can make the font.
					textSize = 1023
				}

				label := material.H1(Theme, "P")
				label.Color = Theme.ContrastFg
				label.TextSize = unit.Sp(float32(textSize) / gtx.Metric.PxPerSp)

				return label.Layout(gtx)
			}),
		)
	})
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
