func sqrt(n int16) int16 {
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


func dlength(a vec2, b vec2) int16 {
	return sqrt((a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y))
}


func pseudoShader() {

	o := vec2{ resolution_x / 2, resolution_y / 2 }

	r := resolution_y / 10
	l := dlength(gl_FragCoord, o)
	c := 255 * r / l

	gl_FragColor = vec4{c, c, c, 255}

}
