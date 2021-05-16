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
// dot(ivec2, ivec2)
// fract(int16) : act as mod 100

func rand100() int16 {
        // based on https://gist.github.com/johansten/3633917
        a := fract(dot(gl_FragCoord, ivec2{ 206, 1245 })) - resolution_y /2
        s := a * (618 + a * a * (-3802 + a * a * 5339))
        t := fract(s * 4735)
        return t
}

func pseudoShader() {
	c := int16(rand100() * 255 / 100)
	gl_FragColor = ivec4{c, c, c, 255}
}
