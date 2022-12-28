// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"bytes"
	_ "embed"
	"image/png"

	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

//go:embed gamer.png
var imageData []byte

var imageOp = func() paint.ImageOp {
	m, err := png.Decode(bytes.NewReader(imageData))
	if err != nil {
		panic(err)
	}
	return paint.NewImageOp(m)
}()

func main() {
	qapp.Render(func(ops *op.Ops) {
		imageOp.Add(ops)
		paint.PaintOp{}.Add(ops)
	})
}
