// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"gioui.org/app"             // app contains Window handling.
	"gioui.org/font/gofont"     // gofont is used for loading the default font.
	"gioui.org/io/system"       // system is used for system events (e.g. closing the window).
	"gioui.org/layout"          // layout is used for layouting widgets.
	"gioui.org/op"              // op is used for recording different operations.
	"gioui.org/widget/material" // material contains material design widgets.
)

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
	// th contains constants for theming.
	th := material.NewTheme(gofont.Collection())
	// ops will be used to encode different operations.
	var ops op.Ops

	counter := 0
	ticker := time.NewTicker(time.Second)

	// listen for events happening on the window.
	for {
		select {
		case <-ticker.C:
			counter++
			w.Invalidate()

		case e := <-w.Events():
			// detect the type of the event.
			switch e := e.(type) {
			// this is sent when the application should re-render.
			case system.FrameEvent:
				// gtx is used to pass around rendering and event information.
				gtx := layout.NewContext(&ops, e)

				layout.Center.Layout(gtx,
					material.H1(th, strconv.Itoa(counter)).Layout,
				)

				// render and handle the operations from the UI.
				e.Frame(gtx.Ops)

			// this is sent when the application is closed.
			case system.DestroyEvent:
				return e.Err
			}
		}

	}
}
