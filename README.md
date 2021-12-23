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
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/singleline/singleline.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/line/line.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/lineongradient/lineongradient.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/wave/wave.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/curvedline/curvedline.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/stepline/stepline.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/smoothstep/smoothstep.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/mix/mix.png)
![](https://github.com/nobuh/toyshader.gba/blob/master/examples/square/square.png)

### Releases

- v0.2: 80x80 screen with 3x2 pixel block version for drawing speed.
- v0.1: 240x160 pixel screen version.

### Build Environemnt

- Ubuntu 20.04 currenly used, but any platform runs TinyGo possible.
- tinygo version 0.21.0 linux/amd64 (using go version go1.17.5 and LLVM version 11.0.0)
- make
- mgba-sdl 0.7.0

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
- pow(x int16, n int16) : x * (x/80)^(n-1)
- mix(x vec3, y vec3, ratio int16) : mix x and y with ratio
- normalizedXY(vec2) : convert native screen x,y position to normalized (0..80) value.
- adjustedRGB(vec3) : convert normalized RGB (0..80) value to adjusted RGB (0..255) value.
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

#### as a ROM

tinygo's GBA target builds a binary as a rom image, so You can run it on any GBA emulator.
However, if you build the binary with default optimization level or '-opt z' optimization level, 
it behaves incorrectly or just hang. Please use opt s or lower optimiazation level.

And if you have Flash Cartridge and its writer, you can burn the binary and run it on real hardware.

#### as a multiboot binary

tinygo's GBA target put the stack on the internal work ram, the heap on the onboard external work ram 
and the program code on the rom area starting from 0x08000000. 

GBA's multiboot mode load progarm at external ram (0x02000000). 
If you want to load the binary on real hardware with multiboot cable, you need to modify the tinygo itself.

pull the tinygo source from github and replace MEMORY directive in the tinygo/target/gameboy-advance.ld 

```
MEMORY {
    ewram   : ORIGIN = 0x02020000, LENGTH = 128K /* upper half of on-board work RAM (2 wait states) */
    iwram   : ORIGIN = 0x03000000, LENGTH = 32K-96 /* in-chip work RAM (faster) */
    rom     : ORIGIN = 0x02000000, LENGTH = 128K  /* lower half of the on-board work RAM */  
}
```

and then build tinygo from the source. 

### Acknowledgements

- [mgba](https://mgba.io/)  
- [TinyGo](https://tinygo.org/) for easy and stable ARM cross compile.
- [META Gameboy Advance Blog](https://remyhax.xyz/posts/gba-blog/) for how to handle Keys.
