// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"

	"gioui.org/f32"
	"gioui.org/font/gofont"     // gofont is used for loading the default font.
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/op/clip"         // clip contains operations for clipping painting area.
	"gioui.org/op/paint"        // paint contains operations for coloring.
	"gioui.org/unit"            // unit contains metric conversion
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
				radius := min(size.X, size.Y) / 2

				shape := clip.Circle{
					Center: f32.Pt(float32(radius), float32(radius)),
					Radius: float32(radius),
				}
				paint.FillShape(gtx.Ops, Theme.ContrastBg, shape.Op(gtx.Ops))

				return layout.Dimensions{
					Size: image.Point{
						X: radius * 2,
						Y: radius * 2,
					},
				}
			}),
			layout.Stacked(func(gtx layout.Context) layout.Dimensions {
				size := gtx.Constraints.Max
				textSize := min(size.X, size.Y) * 3 / 4

				label := material.H1(Theme, "P")
				label.Color = Theme.ContrastFg
				label.TextSize = unit.Px(float32(textSize))
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
