package shader

func shader() {

	o := vec2{ u_resolution.x / 2, u_resolution.y / 2 }

	radius := int16(10)
	l := length(gl_FragCoord, o) / u_resolution.y
	c := cdepth * radius / l

	gl_FragColor = vec3{c, c, c}

}
