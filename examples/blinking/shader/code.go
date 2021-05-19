package shader

func shader() {
	c := abs(sin(u_time * 10)) * 255/100
	gl_FragColor = vec4{c, 0, c, 255}
}
