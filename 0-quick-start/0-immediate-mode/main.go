// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"log"

	"github.com/nsf/termbox-go"
)

func main() {
	var alpha bool
	var beta bool
	var gamma bool
	var delta bool

	Main(func(frame *Frame) {
		Checkbox{Label: "Alpha", Value: &alpha}.Layout(frame)
		Checkbox{Label: "Beta", Value: &beta}.Layout(frame)
		Checkbox{Label: "Gamma", Value: &gamma}.Layout(frame)
		Checkbox{Label: "Delta", Value: &delta}.Layout(frame)
	})
}

/* Example Component */

type Checkbox struct {
	Label string
	Value *bool
}

func (box Checkbox) Layout(frame *Frame) {
	frame.Focus.Input(func() {
		focused := frame.Focus.Active()
		if focused {
			switch frame.Key {
			case termbox.KeyArrowUp:
				frame.Focus.Prev()
			case termbox.KeyArrowDown:
				frame.Focus.Next()
			case termbox.KeyEnter, termbox.KeySpace:
				*box.Value = !*box.Value
			}
		}

		var check = "[ ] "
		if *box.Value {
			check = "[x] "
		}

		if focused {
			frame.Draw.Highlight(check, box.Label)
		} else {
			frame.Draw.Default(check, box.Label)
		}
	})
}

/* Immediate Mode TUI */

// Main starts the main loop of a immediate mode loop.
func Main(fn func(frame *Frame)) {
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	var frame Frame
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventResize:
			ev.Key = 0
			fallthrough
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlC {
				break loop
			}

			frame.Key = ev.Key

			frame.Draw.Invalidate = true
			for frame.Draw.Invalidate {
				frame.Draw.BeforeFrame()

				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				fn(&frame)
				termbox.Flush()

				if frame.Focus.AfterFrame() {
					frame.Draw.Invalidate = true
				}

				frame.Key = 0
			}
		}
	}
}

// Frame contains all information to update and draw the UI.
type Frame struct {
	Input
	Draw  Draw
	Focus Focus
}

// Input contains information about the input key.
type Input struct {
	Key termbox.Key
}

// Focus tracks currently focused items and next focus state.
type Focus struct {
	current int
	active  int
	next    int
}

// Input creates a widget that can be focused.
func (focus *Focus) Input(fn func()) {
	focus.current++
	if focus.active < 0 {
		focus.active = 0
		focus.next = 0
	}
	fn()
}

// Set sets the focus to the current context.
func (focus *Focus) Set() { focus.next = focus.current }

// Prev moves focus backwards.
func (focus *Focus) Prev() { focus.next = focus.active - 1 }

// Next moves focus forward.
func (focus *Focus) Next() { focus.next = focus.active + 1 }

// AfterFrame updates the focus state.
func (focus *Focus) AfterFrame() (updated bool) {
	last := focus.active

	focus.active = focus.next
	if focus.current > 0 {
		if focus.active > focus.current {
			focus.active = focus.active % (focus.current + 1)
		}
		for focus.active < 0 {
			focus.active += (focus.current + 1)
		}
	} else {
		focus.active = -1
	}
	focus.current = -1
	focus.next = focus.active

	return last != focus.active
}

// Active checks whether the current input is in focus.
func (focus *Focus) Active() bool {
	return focus.current == focus.active
}

// Draw contains operations to updating the screen.
type Draw struct {
	Line       int
	Invalidate bool
}

// BeforeFrame resets the draw state.
func (draw *Draw) BeforeFrame() {
	draw.Line = 0
	draw.Invalidate = false
}

// Default draws text with default styling.
func (draw *Draw) Default(args ...interface{}) {
	for x, c := range fmt.Sprint(args...) {
		termbox.SetCell(x, draw.Line, c, termbox.ColorWhite, termbox.ColorBlack)
	}
	draw.Line++
}

// Highlight draws text with inverted styling.
func (draw *Draw) Highlight(args ...interface{}) {
	for x, c := range fmt.Sprint(args...) {
		termbox.SetCell(x, draw.Line, c, termbox.ColorBlack, termbox.ColorWhite)
	}
	draw.Line++
}
