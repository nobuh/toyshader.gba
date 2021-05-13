package main

// Draw a red square on the GameBoy Advance screen.

import (
	"image/color"
	"machine"
)

var display = machine.Display

const CMAX = 255

type vec2 struct {
	x float32
	y float32
}

var resolution = vec2{ 240, 160 }

type vec4 struct {
	r float32
	g float32
	b float32
	a float32
}

var gl_FragColor = vec4{ 0, 0, 0, 0 }


func PseudoShader() {

	gl_FragColor = vec4{ 0.5, 0.5, 0.5, 0.5 }

}


func main() {
	display.Configure()

	for x:= int16(0); x < int16(resolution.x); x++ {
		for y := int16(0); y < int16(resolution.y); y++ {

			PseudoShader()

			r := uint8(gl_FragColor.r * gl_FragColor.a * CMAX)
			g := uint8(gl_FragColor.g * gl_FragColor.a * CMAX)
			b := uint8(gl_FragColor.b * gl_FragColor.a * CMAX)

			display.SetPixel(x, y, color.RGBA{ r, g, b, 0})  // alpha has no effect
		}
	}

	display.Display()
}
