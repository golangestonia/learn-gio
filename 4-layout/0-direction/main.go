// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"gioui.org/font/gofont"     // gofont is used for loading the default font.
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var Theme = material.NewTheme(gofont.Collection())

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		title := material.H1(Theme, "Hello!")
		return layout.Center.Layout(gtx, title.Layout)
	})
}
