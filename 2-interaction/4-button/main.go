// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"gioui.org/app"        // app contains Window handling.
	"gioui.org/gesture"    // gesture contains different gesture event handling.
	"gioui.org/io/event"   // event contains general event information.
	"gioui.org/io/pointer" // pointer contains input/output for mouse and touch screens.
	"gioui.org/io/system"  // system is used for system events (e.g. closing the window).
	"gioui.org/layout"     // layout is used for layouting widgets.
	"gioui.org/op"         // op is used for recording different operations.
	"gioui.org/op/clip"    // clip contains operations for clipping painting area.
	"gioui.org/op/paint"   // paint contains operations for coloring.
)

var button = Button{
	Area: image.Rect(50, 50, 150, 150),
}

func render(ops *op.Ops, queue event.Queue, size image.Point) {
	if button.Render(ops, queue) {
		fmt.Println("clicked")
	}
}

type Button struct {
	Area  image.Rectangle
	click gesture.Click
	color color.NRGBA
}

// Render is a way to implement a button.
func (button *Button) Render(ops *op.Ops, queue event.Queue) (clicked bool) {
	if button.color == (color.NRGBA{}) {
		button.color = color.NRGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xFF}
	}

	// set the area where we want to listen to clicks
	pointer.Rect(button.Area).Add(ops)
	// register click gesture
	button.click.Add(ops)

	// calculate the color of a rectangle based on the click status
	targetColor := color.NRGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xFF}
	if button.click.Hovered() {
		targetColor = color.NRGBA{R: 0x80, G: 0x80, B: 0xA0, A: 0xFF}
	}
	if button.click.Pressed() {
		targetColor = color.NRGBA{R: 0x80, G: 0xA0, B: 0x80, A: 0xFF}
	}

	// animate the color change
	if button.color != targetColor {
		button.color = transition(button.color, targetColor)
		op.InvalidateOp{}.Add(ops)
	}

	// see whether we had a click event
	for _, ev := range button.click.Events(queue) {
		switch ev.Type {
		case gesture.TypeClick:
			button.color = color.NRGBA{R: 0x80, G: 0xFF, B: 0x80, A: 0xFF}
			clicked = true
		}
	}

	// draw the button
	clip.Rect(button.Area).Add(ops)
	paint.ColorOp{Color: button.color}.Add(ops)
	paint.PaintOp{}.Add(ops)

	return clicked
}

func transition(from, to color.NRGBA) color.NRGBA {
	return color.NRGBA{
		R: transitionByte(from.R, to.R),
		G: transitionByte(from.G, to.G),
		B: transitionByte(from.B, to.B),
		A: transitionByte(from.A, to.A),
	}
}

func transitionByte(a, b byte) byte {
	const speed = 2
	delta := int(b) - int(a)
	if delta < -speed {
		delta = -speed
	} else if delta > speed {
		delta = speed
	}
	return byte(int(a) + delta)
}

/* usual setup code */

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
