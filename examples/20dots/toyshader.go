package main

// Draw a red square on the GameBoy Advance screen.

import (
	"image/color"
	"machine"
	"math"
)

var display = machine.Display

type vec2 struct {
	x float32
	y float32
}

type vec4 struct {
	r float32
	g float32
	b float32
	a float32
}

const CMAX = 255
var resolution = vec2{ 240, 160 }
var gl_FragColor = vec4{ 0, 0, 0, 1 }
var gl_FragCoord = vec2{ 0, 0 }

func length(a vec2, b vec2) float64 {
	var l float64
	l = float64((a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y))

	return math.Sqrt(l)
}

/* Your Code Here */

func PseudoShader() {

 	var ncircle float64 = 20.0
 	var orbit_radius float64 = 0.2
 	var circle_radius float64 = 0.01

	var pos vec2 
	pos.x = (gl_FragCoord.x - 0.5 * resolution.x) / resolution.x 
	pos.y = (gl_FragCoord.y - 0.5 * resolution.y) / resolution.x 

	var color = vec4{ 0, 0, 0, 1 }

 	for angle := float64(0); angle < float64(2.0 * 3.14); angle += (2.0 * 3.14 / ncircle) {
		var p vec2
         	p.x = float32(orbit_radius * math.Cos(angle))
         	p.y = float32(orbit_radius * math.Sin(angle))
		len := length(pos, p)

         	color.r += float32(circle_radius / len)
         	color.g += float32(circle_radius / len)
         	color.b += float32(circle_radius / len)
	}

	gl_FragColor.r = color.r
	gl_FragColor.g = color.g
	gl_FragColor.b = color.b
}

/* Until this */


func main() {
	display.Configure()

	for x:= int16(0); x < int16(resolution.x); x++ {
		for y := int16(0); y < int16(resolution.y); y++ {

			gl_FragCoord = vec2{ float32(x), float32(y) }

			PseudoShader()

			r := uint8(gl_FragColor.r * gl_FragColor.a * CMAX)
			g := uint8(gl_FragColor.g * gl_FragColor.a * CMAX)
			b := uint8(gl_FragColor.b * gl_FragColor.a * CMAX)

			display.SetPixel(x, y, color.RGBA{ r, g, b, 0})  // alpha has no effect
		}
	}

	display.Display()
}
