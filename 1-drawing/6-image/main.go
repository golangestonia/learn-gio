// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"

	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
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
		clip.Rect{Max: image.Pt(200, 200)}.Add(ops)
		imageOp.Add(ops)
		paint.PaintOp{}.Add(ops)
	})
}
