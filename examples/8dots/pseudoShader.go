func pseudoShader() {

	const ncircle = int16(8)
	const orbit_radius  = resolution_y / 3
	const dot_radius = int16(30)

	o := vec2{ resolution_x / 2, resolution_y / 2 }

	c := int16(0);
	for i := int16(0); i < 360; i += 360 / ncircle {
		x := orbit_radius * sin(i) / 100 + o.x
		y := orbit_radius * cos(i) / 100 + o.y
		l := length(gl_FragCoord, vec2{ x, y })
		c += 255 * dot_radius / l
	}

	if c > 255 { c = 255 }
	gl_FragColor = vec4{c, c, c, 255}
}
