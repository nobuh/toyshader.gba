package shader

func shader() {
	gray := cdepth / 4
	gl_FragColor = vec3{gray, gray, gray}
}
