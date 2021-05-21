package shader

func slength(a vec2, b vec2) int16 {
	return (a.x - b.x) + (a.y - b.y)
}

func shader() {
	o := vec2{ u_resolution.x / 2, u_resolution.y / 2 }

	r := u_resolution.y / 20
	l := slength(gl_FragCoord, o)
	c := cdepth * r / l

	gl_FragColor = vec3{c, c, c}
}
