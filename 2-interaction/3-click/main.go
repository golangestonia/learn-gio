// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"image"
	"image/color"
	"log"

	"gioui.org/gesture"  // gesture contains different gesture event handling.
	"gioui.org/io/event" // event contains general event information.
	"gioui.org/op"       // op is used for recording different operations.
	"gioui.org/op/clip"  // clip contains operations for clipping painting area.
	"gioui.org/op/paint" // paint contains operations for coloring.

	"github.com/golangestonia/learn-gio/qapp" // qapp contains convenience funcs for this tutorial
)

var buttonClick gesture.Click
var buttonArea = image.Rect(50, 50, 150, 150)
var buttonColor = color.NRGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xFF}

func main() {
	qapp.Input(func(ops *op.Ops, queue event.Queue) {
		// set the area where we want to listen to clicks
		defer clip.Rect(buttonArea).Push(ops).Pop()
		// register click gesture
		buttonClick.Add(ops)

		// calculate the color of a rectangle based on the click status
		targetColor := color.NRGBA{R: 0x40, G: 0x40, B: 0x40, A: 0xFF}
		if buttonClick.Hovered() {
			targetColor = color.NRGBA{R: 0x80, G: 0x80, B: 0xA0, A: 0xFF}
		}
		if buttonClick.Pressed() {
			targetColor = color.NRGBA{R: 0x80, G: 0xA0, B: 0x80, A: 0xFF}
		}

		// animate the color change
		if buttonColor != targetColor {
			buttonColor = transition(buttonColor, targetColor)
			op.InvalidateOp{}.Add(ops)
		}

		// see whether we had a click event
		for _, ev := range buttonClick.Events(queue) {
			switch ev.Type {
			case gesture.TypeClick:
				buttonColor = color.NRGBA{R: 0x80, G: 0xFF, B: 0x80, A: 0xFF}
				log.Println("clicked")
			}
		}

		// draw the button
		clip.Rect(buttonArea).Push(ops).Pop()
		paint.ColorOp{Color: buttonColor}.Add(ops)
		paint.PaintOp{}.Add(ops)
	})
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
