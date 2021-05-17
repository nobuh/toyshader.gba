func pseudoShader() {

	o := vec2{ resolution_x / 2, resolution_y / 2 }

	radius := int16(30)
	l := length(gl_FragCoord, o)

	c := int16(0)
	if l < radius {
		c = 255
	} else {
		c = 255 * radius / l
	}

	gl_FragColor = vec4{c, c, c, 255}

}
