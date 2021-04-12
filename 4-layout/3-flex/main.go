// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"image/color"

	"gioui.org/font/gofont"     // gofont is used for loading the default font.
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/text"            // text contains constants for text layouting.
	"gioui.org/unit"            // unit contains metric conversion
	"gioui.org/widget"          // widget contains state for different widgets
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var Theme = material.NewTheme(gofont.Collection())

var editor widget.Editor

var inset = layout.UniformInset(unit.Dp(8))
var border = widget.Border{
	Color: color.NRGBA{R: 0x88, G: 0x88, B: 0x88, A: 0xFF},
	Width: unit.Px(1),
}

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		return inset.Layout(gtx,
			func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis:      layout.Vertical,
					Alignment: layout.Middle,
				}.Layout(gtx,
					layout.Rigid(Center(material.H2(Theme, "Header")).Layout),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return border.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
							return inset.Layout(gtx, material.Editor(Theme, &editor, "").Layout)
						})
					}),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						line, col := editor.CaretPos()
						s := fmt.Sprintf("line:%d col:%d", line, col)
						return Center(material.Body1(Theme, s)).Layout(gtx)
					}),
				)
			})
	})
}

func Center(label material.LabelStyle) material.LabelStyle {
	label.Alignment = text.Middle
	return label
}
