# toyshader.gba

Toy Shader in TinyGo for Game Boy Advance.

![](https://github.com/nobuh/toyshader.gba/blob/master/examples/gray.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/slope.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/rings.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/diamond.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/blackhole.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/dot.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/8dots.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/rand.png)

![](https://github.com/nobuh/toyshader.gba/blob/master/examples/blinking.gif)

### Build Environemnt

- Ubuntu 20.04 currenly used, but any platform runs TinyGo possible.
- Tiny Go version 0.18.0 linux/amd64 (using go version go1.16.4 and LLVM version 11.0.0)
- make
- VisualBoyAdvance version 1.8.0

### How to Build examples.

- git clone https://github.com/nobuh/toyshader.gba.git
- cd toyshader.gba/examples/[any demo you want to build]/
- make

### How to run this demo on real hardware

I'm currently using Visual Boy Advance. Because it can load and run ROM cart memory (0x08000000) without Nintendo's ROM headers.
It seems that the Visual Boy Advance can emulate the 3rd party Flash Cart Drive which has no Nintendo headers, so we don'nt need create ROM headers.  

If you want to run this demo on real hardware or mGBA, pleae check devkitPro or similer tools.

### Usage 

from examples/gray/pseudoShader.go

```
// Write Your Shader in pseudoShader()
//
// resolution_x = 240
// resolution_y = 160
//
// all variable type should be int16
// vec2 has { x,y }
// vec4 has { r,g,b,a }
//
// R,G,B,A = 0..255, 0..255, 0..255, 0..255
//
// sin(degree) sin with radius = 100
// cos(degree) sin with radius = 100
// length(ivec2, ivec2) : euclidean distance ^2
// dot(ivec2, ivec2) 
// fract(int16) : act as mod 100

func pseudoShader() {
	gray := int16(255/2)
	half_transparent := int16(255/2)
	gl_FragColor = ivec4{gray, gray, gray, half_transparent}
}
```

append your pseudoShader.go file to the toyshader.go and build it by tinygo compiler.

