// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image/color"

	"gioui.org/font/gofont"     // gofont is used for loading the default font.
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/text"            // text contains constants for text layouting.
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var Theme = material.NewTheme(gofont.Collection())

func main() { qapp.Layout(Layout) }

// Layout handles rendering and input.
func Layout(gtx layout.Context) layout.Dimensions {
	return Title(Theme, "Hello, Quick!").Layout(gtx)
}

// Title creates a center aligned H1.
func Title(th *material.Theme, caption string) material.LabelStyle {
	l := material.H1(th, caption)
	l.Color = color.NRGBA{R: 127, G: 0, B: 0, A: 255}
	l.Alignment = text.Middle
	return l
}
