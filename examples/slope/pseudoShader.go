func slength(a vec2, b vec2) int16 {
	return (a.x - b.x) + (a.y - b.y)
}

func pseudoShader() {
	o := vec2{ resolution_x / 2, resolution_y / 2 }

	r := resolution_y / 20
	l := slength(gl_FragCoord, o)
	c := 255 * r / l

	gl_FragColor = vec4{c, c, c, 255}
}
