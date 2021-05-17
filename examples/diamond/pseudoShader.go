func difflength(a vec2, b vec2) int16 {
	var x, y int16

	if a.x > b.x {
		x = a.x - b.x
	} else {
		x = b.x - a.x
	}

	if a.y > b.y {
		y = a.y - b.y
	} else {
		y = b.y - a.y
	}

	return x + y
}


func pseudoShader() {

	o := vec2{ resolution_x / 2, resolution_y / 2 }

	radius := resolution_y / 20
	l := difflength(gl_FragCoord, o)
	if l < radius {
		l = radius
	}
	c := 255 * radius / l

	gl_FragColor = vec4{c, c, c, 255}

}
