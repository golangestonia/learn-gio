// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
	"log"
	"os"

	"gioui.org/app"       // app contains Window handling.
	"gioui.org/f32"       // f32 contains float32 points.
	"gioui.org/io/event"  // event contains general event information.
	"gioui.org/io/key"    // key contains input/output for keyboards.
	"gioui.org/io/system" // system is used for system events (e.g. closing the window).
	"gioui.org/layout"    // layout is used for layouting widgets.
	"gioui.org/op"        // op is used for recording different operations.
	"gioui.org/op/paint"  // paint contains operations for coloring.
)

var location = f32.Pt(300, 300)

const speed = 50

func render(ops *op.Ops, queue event.Queue, size image.Point) {
	// keep the focus, since only one thing can
	key.FocusOp{Tag: &location}.Add(ops)
	// register tag &location as reading input
	key.InputOp{Tag: &location}.Add(ops)

	// read events from input event queue
	for _, ev := range queue.Events(&location) {
		// figure out, which event it was
		switch ev := ev.(type) {
		case key.Event:
			if ev.State == key.Press {
				switch ev.Name {
				case key.NameLeftArrow:
					location.X -= speed
				case key.NameUpArrow:
					location.Y -= speed
				case key.NameRightArrow:
					location.X += speed
				case key.NameDownArrow:
					location.Y += speed
				}
			}
		}
	}

	op.Offset(location).Add(ops)
	imageOp.Add(ops)
	paint.PaintOp{}.Add(ops)
}

/* usual setup code */

//go:embed neutral.png
var imageData []byte

var imageOp = func() paint.ImageOp {
	m, err := png.Decode(bytes.NewReader(imageData))
	if err != nil {
		panic(err)
	}
	return paint.NewImageOp(m)
}()

func main() {
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	// ops will be used to encode different operations.
	var ops op.Ops

	// listen for events happening on the window.
	for e := range w.Events() {
		// detect the type of the event.
		switch e := e.(type) {
		// this is sent when the application should re-render.
		case system.FrameEvent:
			// gtx is used to pass around rendering and event information.
			gtx := layout.NewContext(&ops, e)
			// render content
			render(gtx.Ops, gtx.Queue, gtx.Constraints.Max)
			// render and handle the operations from the UI.
			e.Frame(gtx.Ops)

		// this is sent when the application is closed.
		case system.DestroyEvent:
			return e.Err
		}
	}

	return nil
}
