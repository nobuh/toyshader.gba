package shader

func shader() {
	y := smoothstep(10, 0, abs(gl_FragCoord.y - gl_FragCoord.x))

	color := vec3{ 0, cdepth * y / normalize, 0 }

	gl_FragColor = color
}
