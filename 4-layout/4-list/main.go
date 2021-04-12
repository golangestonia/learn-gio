// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"

	"gioui.org/font/gofont"     // gofont is used for loading the default font.
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/widget/material" // material contains material design widgets.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var Theme = material.NewTheme(gofont.Collection())

var list = layout.List{
	Axis: layout.Vertical,
}

var items = createItems(1000)

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		return list.Layout(gtx, len(items),
			func(gtx layout.Context, index int) layout.Dimensions {
				return material.Body1(Theme, items[index]).Layout(gtx)
			})
	})
}

func createItems(n int) []string {
	xs := []string{}
	for i := 0; i < n; i++ {
		xs = append(xs, fmt.Sprintf("%08x", i))
	}
	return xs
}
