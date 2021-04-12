// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/font/gofont"     // gofont is used for loading the default font.
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/op/paint"        // paint contains operations for coloring.
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var Theme = func() *material.Theme {
	theme := material.NewTheme(gofont.Collection())
	theme.Palette = Invert(theme.Palette)
	return theme
}()

func Invert(pal material.Palette) material.Palette {
	pal.Fg, pal.Bg = pal.Bg, pal.Fg
	pal.ContrastFg, pal.ContrastBg = pal.ContrastBg, pal.ContrastFg
	return pal
}

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		paint.ColorOp{Color: Theme.Bg}.Add(gtx.Ops)
		paint.PaintOp{}.Add(gtx.Ops)

		return material.H1(Theme, "Hello, Quick!").Layout(gtx)
	})
}
