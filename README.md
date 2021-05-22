# toyshader.gba

Toy Shader in TinyGo for Game Boy Advance.

![](https://github.com/nobuh/toyshader.gba/blob/master/examples/gray/gray.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/slope/slope.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/rings/rings.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/diamond/diamond.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/blackhole/blackhole.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/dot/dot.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/8dots/8dots.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/rand/rand.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/blinking/blinking.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/gradient/gradient.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/blinkinggradient/blinkinggradient.png)

### Releases

- v0.2: 80x80 screen with 3x2 pixel block version for drawing speed.
- v0.1: 240x160 pixel screen version.


### Build Environemnt

- Ubuntu 20.04 currenly used, but any platform runs TinyGo possible.
- Tiny Go version 0.18.0 linux/amd64 (using go version go1.16.4 and LLVM version 11.0.0)
- make
- VisualBoyAdvance version 1.8.0

### Build examples.

- export GO111MODULE="off"
- cd $GOPATH/src
- git clone https://github.com/nobuh/toyshader.gba.git
- cd toyshader.gba/examples/[any demo you want to build]/
- make

### Shader Values and Functions

all variable type is int16

R,G,B = { 0..255, 0..255, 0..255 }

- u_resolution = { 240, 160 }
- u_time = 0..35
- u_mouse = { 0..240, 0..160 }
- cdepth = 255 // color depth
- vec2 has { x,y } 
- vec3 has { r,g,b }
- sin() / cos() : radius = 80 and degree by 10
- length(vec2, vec2) : euclidean distance ^2
- dot(vec2, vec2) 
- fract(int16) : att like mod 80
- shader.TurnPageByKey() : need to be called before Run() and you can turn screen by any key
- shader.Run() : main loop

### Build your shader

- export GO111MODULE="off"
- cd $GOPATH/src
- git clone https://github.com/nobuh/toyshader.gba.git
- cd toyshader.gba
- write your shader in function shader() on the shader/code.go
- if you want to use TurnPageByKey() mode, put it on main.go before Run()
- make

### How to run this demo on real hardware

I'm currently using Visual Boy Advance. Because it can load and run ROM cart memory (0x08000000) without Nintendo's ROM headers.
It seems that the Visual Boy Advance can emulate the 3rd party Flash Cart Drive which has no Nintendo headers, so we don'nt need create ROM headers.  

If you want to run this demo on real hardware or mGBA, pleae check devkitPro or similer tools.

### Acknowledgements

- [VisualBoyAdvance](https://board.vba-m.com/) I 'd like to use raw binary as ROM, without fake Nintendo's logo.
- [TinyGo](https://tinygo.org/) for easy and stable ARM cross compile.
- [META Gameboy Advance Blog](https://remyhax.xyz/posts/gba-blog/) for how to handle Keys.
