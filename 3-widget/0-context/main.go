// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"
	"log"

	"gioui.org/gesture"  // gesture contains different gesture event handling.
	"gioui.org/layout"   // layout is used for layouting widgets.
	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

/*
// Context carries the state needed by almost all layouts and widgets.
type Context struct {
	// Constraints track the constraints for the active widget or layout.
	Constraints Constraints

	Metric unit.Metric
	// By convention, a nil Queue is a signal to widgets to draw themselves
	// in a disabled state.
	Queue event.Queue
	// Now is the animation time.
	Now time.Time

	*op.Ops
}
*/

var button = Button{
	Area: image.Rect(50, 50, 150, 150),
}

func main() {
	qapp.Layout(func(gtx layout.Context) layout.Dimensions {
		if button.Layout(gtx) {
			log.Println("clicked")
		}

		return layout.Dimensions{}
	})
}

type Button struct {
	Area  image.Rectangle
	click gesture.Click
	color color.NRGBA
}

// Layout is a way to implement a button.
func (button *Button) Layout(gtx layout.Context) (clicked bool) {
	if button.color == (color.NRGBA{}) {
		button.color = color.NRGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xFF}
	}

	// set the area where we want to listen to clicks
	defer clip.Rect(button.Area).Push(gtx.Ops).Pop()
	// register click gesture
	button.click.Add(gtx.Ops)

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
		// TODO: this should use gtx.Now for color changes
		button.color = transition(button.color, targetColor)
		op.InvalidateOp{}.Add(gtx.Ops)
	}

	// see whether we had a click event
	for _, ev := range button.click.Events(gtx.Queue) {
		switch ev.Type {
		case gesture.TypeClick:
			button.color = color.NRGBA{R: 0x80, G: 0xFF, B: 0x80, A: 0xFF}
			clicked = true
		}
	}

	// draw the button
	defer clip.Rect(button.Area).Push(gtx.Ops).Pop()
	paint.ColorOp{Color: button.color}.Add(gtx.Ops)
	paint.PaintOp{}.Add(gtx.Ops)

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
