package shader

import (
	"image/color"
	"machine"
	"math"
	"runtime/interrupt"
	"runtime/volatile"
	"unsafe"
)


// ========== Global Components ==========

var display = machine.Display

type vec2 struct {
	x int16
	y int16
}

type vec3 struct {
	r int16
	g int16
	b int16
}

const cdepth   = int16(255) // color depth

var gl_FragColor = vec3{ 0, 0, 0 } // shoud be 0..cdepth
var gl_FragCoord = vec2{ 0, 0 }	// shoud be 0..norm 

var sinCache[36]int16 // radius = norm 
var cosCache[36]int16 // degree by 10, 0..350

var u_time = int16(0) // pseudo time value 0..35
var u_resolution = vec2{ 240, 160 }
var u_mouse = vec2{ u_resolution.x / 2, u_resolution.y / 2 }

var KeyEnabled uint8 = 0 // Key is default disabled

// from https://remyhax.xyz/posts/gba-blog/
var regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))
var regKEYPAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000130)))

var (
	//KeyCodes
	keyDOWN      = uint16(895)
	keyUP        = uint16(959)
	keyLEFT      = uint16(991)
	keyRIGHT     = uint16(1007)
	keyLSHOULDER = uint16(511)
	keyRSHOULDER = uint16(767)
	keyA         = uint16(1022)
	keyB         = uint16(1021)
	keySTART     = uint16(1015)
	keySELECT    = uint16(1019)
)

// ========== Shader Functions ==========

func TurnPageByKey() {
	KeyEnabled = uint8(1)
}


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
	return n - n / 80
}


func abs(n int16) int16 {
	if n < 0 { n = - n }
	return n
}


func Run() {
	// init sin cos cache
	for i := int16(0); i < 36; i++ {
		sinCache[i] = int16(80.0 * math.Sin(2.0 * 3.14 * float64(i) / 36))
		cosCache[i] = int16(80.0 * math.Cos(2.0 * 3.14 * float64(i) / 36))
	}

	display.Configure()
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	// from https://remyhax.xyz/posts/gba-blog/
	interrupt.New(machine.IRQ_VBLANK, update).Enable()
	draw()	// first piciture
	for { } // prevent exit
}

func update_time() {
	u_time++
	if u_time > 35 { u_time = 0 }
}


func update(interrupt.Interrupt) {
	switch keyValue := regKEYPAD.Get(); keyValue {
	case keyDOWN:
		u_mouse.y += u_resolution.y * 10
		update_time()
		draw()
	case keyUP:
		u_mouse.y -= u_resolution.y * 10
		update_time()
		draw()
	case keyLEFT:
		u_mouse.x -= u_resolution.x * 10
		update_time()
		draw()
	case keyRIGHT:
		u_mouse.x += u_resolution.x * 10
		update_time()
		draw()
	case keyA:
		update_time()
		draw()
	case keyB:
		update_time()
		draw()
	default:
		if KeyEnabled == 0 {
			update_time()
			draw()
		}
	}
}


func draw() {

	for y := int16(0); y < u_resolution.y; y += 2 {
		for x:= int16(0); x < u_resolution.x; x += 3 {

			gl_FragCoord = vec2{ x, y }

			shader() // defined in user code

                        r := uint8(gl_FragColor.r)
                        g := uint8(gl_FragColor.g)
                        b := uint8(gl_FragColor.b)

                        display.SetPixel(x, y, color.RGBA{ r, g, b, 0})   // alpha has no effect
                        display.SetPixel(x+1, y, color.RGBA{ r, g, b, 0})
                        display.SetPixel(x+2, y, color.RGBA{ r, g, b, 0})
                        display.SetPixel(x, y+1, color.RGBA{ r, g, b, 0})
                        display.SetPixel(x+1, y+1, color.RGBA{ r, g, b, 0})
                        display.SetPixel(x+2, y+1, color.RGBA{ r, g, b, 0})
                }
	}
}
