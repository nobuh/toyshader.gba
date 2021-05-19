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

var sinCache[36]int16 // radius = 100
var cosCache[36]int16 // degree by 10, 0..350
var u_time = int16(0) // pseudo time value 0..35

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
	return n - n / 100
}


func abs(n int16) int16 {
	if n < 0 { n = - n }
	return n
}


func Run() {
	// init sin cos cache
	for i := int16(0); i < 36; i++ {
		sinCache[i] = int16(100 * math.Sin(2.0 * 3.14 * float64(i) / 36))
		cosCache[i] = int16(100 * math.Cos(2.0 * 3.14 * float64(i) / 36))
	}

	display.Configure()
	regDISPSTAT.SetBits(1<<3 | 1<<4)

	if KeyEnabled > 0 {
		// from https://remyhax.xyz/posts/gba-blog/
		interrupt.New(machine.IRQ_VBLANK, update).Enable()
		for { } // prevent exit
	} else {
		for {
			update_time()
			draw()
		}
	}
}

func update_time() {
	u_time++
	if u_time > 35 { u_time = 0 }
}


func update(interrupt.Interrupt) {
	switch keyValue := regKEYPAD.Get(); keyValue {
	case keyDOWN:
		update_time()
		draw()
	case keyUP:
		update_time()
		draw()
	case keyLEFT:
		update_time()
		draw()
	case keyRIGHT:
		update_time()
		draw()
	case keyA:
		update_time()
		draw()
	case keyB:
		update_time()
		draw()
	}
}


func draw() {

	for y := int16(0); y < resolution_y; y++ {
		for x:= int16(0); x < resolution_x; x++ {

			gl_FragCoord = vec2{ x, y }

			shader() // defined in user code

                        // gl_FragColor * gl_FragColor can be over int16
                        // so split /255 to /5 and /51
                        r := uint8(gl_FragColor.r * (gl_FragColor.a / 5) / 51)
                        g := uint8(gl_FragColor.g * (gl_FragColor.a / 5) / 51)
                        b := uint8(gl_FragColor.b * (gl_FragColor.a / 5) / 51)

                        display.SetPixel(x, y, color.RGBA{ r, g, b, 0})  // alpha has no effect
                }
	}
}
