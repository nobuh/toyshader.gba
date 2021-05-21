package shader

func shader() {

	o := vec2{ u_resolution.x / 2, u_resolution.y / 2 }

	radius := int16(30)
	l := length(gl_FragCoord, o)

	c := int16(0)
	if l < radius {
		c = cdepth
	} else {
		c = cdepth * radius / l
	}

	gl_FragColor = vec3{c, c, c}

}
