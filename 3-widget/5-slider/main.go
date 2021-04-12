// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/font/gofont"     // gofont is used for loading the default font.
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/widget"          // widget contains state for different widgets
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var Theme = material.NewTheme(gofont.Collection())

var value widget.Float

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		return material.Slider(Theme, &value, 0, 10).Layout(gtx)
	})
}
