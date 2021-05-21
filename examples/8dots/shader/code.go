package shader

func shader() {

	const ncircle = int16(8)
	const dot_radius = int16(30)
	var orbit_radius  = u_resolution.y / 3

	o := vec2{ u_resolution.x / 2, u_resolution.y / 2 }

	c := int16(0);
	for i := int16(0); i < 360; i += 360 / ncircle {
		x := orbit_radius * sin(i) / 100 + o.x
		y := orbit_radius * cos(i) / 100 + o.y
		l := length(gl_FragCoord, vec2{ x, y })
		c += cdepth * dot_radius / l
	}

	if c > cdepth { c = cdepth }
	gl_FragColor = vec3{c, c, c}
}
