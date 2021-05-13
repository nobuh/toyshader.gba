package main

import (
	"image/color"
	"machine"
)

var display = machine.Display

type ivec2 struct {
	x int16
	y int16
}

type ivec4 struct {
	r int16
	g int16
	b int16
	a int16
}

const CMAX int16 = 255
var resolution = ivec2{ 240, 160 }
var gl_FragColor = ivec4{ 0, 0, 0, 255 }
var gl_FragCoord = ivec2{ 0, 0 }

func length(a ivec2, b ivec2) int16 {
	return (a.x - b.x) + (a.y - b.y)
}

/* Your Code Here */

func PseudoShader() {

	o := ivec2{ resolution.x / 2, resolution.y / 2 }

	r := resolution.y / 20 // circle size is 1/20 of y axis
	l := length(gl_FragCoord, o)
	c := CMAX * r / l

	gl_FragColor = ivec4{c, c, c, CMAX}
}

/* Until this */


func main() {
	display.Configure()

	for x:= int16(0); x < resolution.x; x++ {
		for y := int16(0); y < resolution.y; y++ {

			gl_FragCoord = ivec2{ x, y }

			PseudoShader()

			r := uint8(gl_FragColor.r * gl_FragColor.a / CMAX)
			g := uint8(gl_FragColor.g * gl_FragColor.a / CMAX)
			b := uint8(gl_FragColor.b * gl_FragColor.a / CMAX)

			display.SetPixel(x, y, color.RGBA{ r, g, b, 0})  // alpha has no effect
		}
	}

	display.Display()
}
