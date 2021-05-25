package shader

func plot(st vec2) int16 {
	return smoothstep(5, 0, abs(st.y - st.x))
}


func shader() {
	st := vec2{ normalize * gl_FragCoord.x / u_resolution.x, normalize * gl_FragCoord.y / u_resolution.y }
	y := plot(st)
	color := vec3{ 0, cdepth * y / normalize, 0 }
	gl_FragColor = color
}
