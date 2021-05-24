package shader

func plot(st vec2) int16 {
	return smoothstep(5, 0, abs(st.y - st.x))
}


func shader() {
	st := vec2{ normalize * gl_FragCoord.x / u_resolution.x, normalize * gl_FragCoord.y / u_resolution.y }

	y := st.x
	color := vec3{ y, y, y }

	pct := plot(st)
	color.r = (normalize - pct) * color.r
	color.g = (normalize - pct) * color.g + pct * normalize
	color.b = (normalize - pct) * color.b

	gl_FragColor = color
}
