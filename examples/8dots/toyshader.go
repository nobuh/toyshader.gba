package main

import (
	"image/color"
	"machine"
	"math"
)

var display = machine.Display

const resolution_x = int16(240)
const resolution_y = int16(160)

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

var gl_FragColor = ivec4{ 0, 0, 0, 255 }
var gl_FragCoord = ivec2{ 0, 0 }

// integer radius 100 per 10 degree cache
var sin100[36]int16
var cos100[36]int16

func distanceFromO(v ivec2) int16 {
	return v.x * v.x + v.y * v.y
}


func length(a ivec2, b ivec2) int16 {
	return distanceFromO(ivec2{a.x - b.x, a.y - b.y})
}


func pseudoShader() {
	// R,G,B = 0..255, 0..255, 0..255

	const ncircle = int16(8)
	const orbit_radius  = resolution_y / 3
	const dot_radius = int16(30)

	o := ivec2{ resolution_x / 2, resolution_y / 2 }

	c := int16(0);
	for i := int16(0); i < 360; i += 360 / ncircle {
		x := orbit_radius * sin100[i / 10] / 100 + o.x
		y := orbit_radius * cos100[i / 10] / 100 + o.y
		l := length(gl_FragCoord, ivec2{ x, y })
		c += 255 * dot_radius / l
	}

	if c > 255 {
		c = 255
	}
	gl_FragColor = ivec4{c, c, c, 255}

}


func main() {
	// init sin cos cache
	for i := int16(0); i < 36; i++ {
		sin100[i] = int16(100 * math.Sin(2.0 * 3.14 * float64(i) / 36))
		cos100[i] = int16(100 * math.Cos(2.0 * 3.14 * float64(i) / 36))
	}

	display.Configure()

	for y := int16(0); y < resolution_y; y++ {
		for x:= int16(0); x < resolution_x; x++ {

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
