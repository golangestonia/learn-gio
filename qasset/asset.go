// SPDX-License-Identifier: Unlicense OR MIT

package qasset

import (
	"bytes"
	_ "embed"
	"image"
	"image/png"
)

//go:embed neutral.png
var neutralData []byte

var Neutral = func() image.Image {
	m, err := png.Decode(bytes.NewReader(neutralData))
	if err != nil {
		panic(err)
	}
	return m
}()

//go:embed gamer.png
var gamerData []byte

var Gamer = func() image.Image {
	m, err := png.Decode(bytes.NewReader(gamerData))
	if err != nil {
		panic(err)
	}
	return m
}()
