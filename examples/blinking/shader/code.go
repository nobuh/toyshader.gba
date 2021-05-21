package shader

func shader() {
	c := abs(sin(u_time * 10)) * cdepth / 100
	gl_FragColor = vec3{c, 0, c}
}
