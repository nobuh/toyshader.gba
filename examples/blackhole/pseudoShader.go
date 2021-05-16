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

	o := ivec2{ resolution_x / 2, resolution_y / 2 }

	radius := int16(10)
	l := ilength(gl_FragCoord, o) / resolution_y
	c := 255 * radius / l

	gl_FragColor = ivec4{c, c, c, 255}

}
