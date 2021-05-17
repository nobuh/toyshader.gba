func pseudoShader() {

	o := vec2{ resolution_x / 2, resolution_y / 2 }

	radius := int16(10)
	l := length(gl_FragCoord, o) / resolution_y
	c := 255 * radius / l

	gl_FragColor = vec4{c, c, c, 255}

}
