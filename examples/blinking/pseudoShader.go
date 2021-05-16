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
	c := abs(sin100(u_time * 10)) * 255/100
	gl_FragColor = ivec4{c, 0, c, 255}
}
