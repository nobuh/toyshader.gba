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

var resolution = ivec2{ 240, 160 }
var gl_FragColor = ivec4{ 0, 0, 0, 255 }
var gl_FragCoord = ivec2{ 0, 0 }


func length(a ivec2, b ivec2) int16 {
	var x, y int16

	if a.x > b.x {
		x = a.x - b.x
	} else {
		x = b.x - a.x
	}

	if a.y > b.y {
		y = a.y - b.y
	} else {
		y = b.y - a.y
	}

	return x + y
}


func pseudoShader() {
	// R,G,B = 0..255, 0..255, 0..255

	o := ivec2{ resolution.x / 2, resolution.y / 2 }

	radius := resolution.y / 20
	l := length(gl_FragCoord, o)
	if l < radius {
		l = radius
	}
	c := 255 * radius / l

	gl_FragColor = ivec4{c, c, c, 255}

}


func main() {
	display.Configure()

	for y := int16(0); y < resolution.y; y++ {
		for x:= int16(0); x < resolution.x; x++ {

			gl_FragCoord = ivec2{ x, y }

			pseudoShader()

			r := uint8(gl_FragColor.r * gl_FragColor.a / 255)
			g := uint8(gl_FragColor.g * gl_FragColor.a / 255)
			b := uint8(gl_FragColor.b * gl_FragColor.a / 255)

			display.SetPixel(x, y, color.RGBA{ r, g, b, 0})  // alpha has no effect
		}
	}

	display.Display()
}
