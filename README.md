# toyshader.gba

Toy Shader in TinyGo for Game Boy Advance.

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
// ivec2 has x,y 
// ivec4 has R,G,B,A
//
// R,G,B,A = 0..255, 0..255, 0..255, 0..255
//
// sin100 : sin(degree) with radius = 100
// cos100 : cos(degree) with radius = 100
// ilength(ivec2, ivec2) : euclidean distance ^2

func pseudoShader() {
	gray := int16(255/2)
	half_transparent := int16(255/2)
	gl_FragColor = ivec4{gray, gray, gray, half_transparent}
}
```

append your pseudoShader.go file to the toyshader.go and build it by tinygo compiler.

