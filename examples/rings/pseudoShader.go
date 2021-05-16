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

func isqrt(n int16) int16 {
	x := n
	num_iterations := int16(5)
	guess := int16(0)

	for i := int16(0); i < num_iterations; i++ {
		guess = (x + (n / x) ) / 2 * 100
		if (guess - x) < 1 {
			break
		} else {
			x = guess
		}
	}
	return int16(guess)
}


func length(a ivec2, b ivec2) int16 {
	return isqrt((a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y))
}


func pseudoShader() {

	o := ivec2{ resolution_x / 2, resolution_y / 2 }

	r := resolution_y / 10
	l := length(gl_FragCoord, o)
	c := 255 * r / l

	gl_FragColor = ivec4{c, c, c, 255}

}
