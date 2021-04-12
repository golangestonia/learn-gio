// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"log"

	"gioui.org/font/gofont" // gofont is used for loading the default font.
	"gioui.org/layout"      // layout is used for layouting widgets.
	"gioui.org/unit"
	"gioui.org/widget"          // widget contains state for different widgets
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var Theme = material.NewTheme(gofont.Collection())

var click widget.Clickable

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		for click.Clicked() {
			log.Println("Click")
		}
		// layout.Inset{}
		return layout.UniformInset(unit.Dp(32)).Layout(gtx,
			material.Button(Theme, &click, "Click").Layout,
		)
	})
}
