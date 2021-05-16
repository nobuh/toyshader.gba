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

// radius 100, sin/cos each 10 degrees
var sin100Cache[36]int16
var cos100Cache[36]int16


func sin100(degree int16) int16 {
	return sin100Cache[degree / 10]
}


func cos100(degree int16) int16 {
	return cos100Cache[degree / 10]
}


func distanceFromO(v ivec2) int16 {
	l := int32(v.x * v.x + v.y * v.y)
	if l > 32767 { l = 32767 }
	return int16(l)
}


func ilength(a ivec2, b ivec2) int16 {
	return distanceFromO(ivec2{a.x - b.x, a.y - b.y})
}


func dot(a ivec2, b ivec2) int16 {
	return a.x * b.x + a.y + b.y
}


func fract(n int16) int16 {
	return n - n / 100
}


func main() {
	// init sin cos cache
	for i := int16(0); i < 36; i++ {
		sin100Cache[i] = int16(100 * math.Sin(2.0 * 3.14 * float64(i) / 36))
		cos100Cache[i] = int16(100 * math.Cos(2.0 * 3.14 * float64(i) / 36))
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
