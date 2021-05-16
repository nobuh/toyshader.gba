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

	const ncircle = int16(8)
	const orbit_radius  = resolution_y / 3
	const dot_radius = int16(30)

	o := ivec2{ resolution_x / 2, resolution_y / 2 }

	c := int16(0);
	for i := int16(0); i < 360; i += 360 / ncircle {
		x := orbit_radius * sin100(i) / 100 + o.x
		y := orbit_radius * cos100(i) / 100 + o.y
		l := ilength(gl_FragCoord, ivec2{ x, y })
		c += 255 * dot_radius / l
	}

	if c > 255 { c = 255 }
	gl_FragColor = ivec4{c, c, c, 255}
}
