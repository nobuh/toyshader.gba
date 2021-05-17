package main

import (
	"image/color"
	"machine"
	"math"
)

var display = machine.Display

const resolution_x = int16(240)
const resolution_y = int16(160)

type vec2 struct {
	x int16
	y int16
}

type vec4 struct {
	r int16
	g int16
	b int16
	a int16
}

var gl_FragColor = vec4{ 0, 0, 0, 255 }
var gl_FragCoord = vec2{ 0, 0 }

// radius 100, sin/cos each 10 degrees
var sinCache[36]int16
var cosCache[36]int16

var u_time = int16(0) // pseudo time value 0..35

func sin(degree int16) int16 {
	degree = degree - degree / 360
	return sinCache[degree / 10]
}


func cos(degree int16) int16 {
	degree = degree - degree / 360
	return cosCache[degree / 10]
}


func distanceFromO(v vec2) int16 {
	l := int32(v.x * v.x + v.y * v.y)
	if l > 32767 { l = 32767 }
	return int16(l)
}


func length(a vec2, b vec2) int16 {
	return distanceFromO(vec2{a.x - b.x, a.y - b.y})
}


func dot(a vec2, b vec2) int16 {
	return a.x * b.x + a.y + b.y
}


func fract(n int16) int16 {
	return n - n / 100
}


func abs(n int16) int16 {
	if n < 0 { n = - n }
	return n
}



func main() {
	// init sin cos cache
	for i := int16(0); i < 36; i++ {
		sinCache[i] = int16(100 * math.Sin(2.0 * 3.14 * float64(i) / 36))
		cosCache[i] = int16(100 * math.Cos(2.0 * 3.14 * float64(i) / 36))
	}

	display.Configure()

	for {
		// update current u_time
		u_time++
		if u_time > 35 { u_time = 0 }

		for y := int16(0); y < resolution_y; y++ {
			for x:= int16(0); x < resolution_x; x++ {

				gl_FragCoord = vec2{ x, y }

				pseudoShader()

				r := uint8(gl_FragColor.r * gl_FragColor.a / 255)
				g := uint8(gl_FragColor.g * gl_FragColor.a / 255)
				b := uint8(gl_FragColor.b * gl_FragColor.a / 255)

				display.SetPixel(x, y, color.RGBA{ r, g, b, 0})  // alpha has no effect
			}
		}

		display.Display()
	}
}
func pseudoShader() {
	// if 255 * 240, the value over int16, so divide 255 to 51 * 5 
	st := vec2{ abs(cos(u_time * 10)) * 51/100 * gl_FragCoord.x / resolution_x * 5,
		    abs(sin(u_time * 10)) * 51/100 * gl_FragCoord.y / resolution_y * 5 }
	gl_FragColor = vec4{ 0, st.x, st.y, 255}
}
