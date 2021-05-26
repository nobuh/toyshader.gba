package shader

func plot(st vec2) int16 {
	return smoothstep(5, 0, abs(st.y - st.x)) - smoothstep(5, 0, abs(st.y - st.x))
}

func shader() {
	st := normalizedXY(gl_FragCoord)

	// gradient
	y := sin(st.x * u_time)
	color := vec3{ y, y, y }

	// line
	pct := plot(st)
	color = vec3 { (normalize - pct) * color.r / normalize,
		       (normalize - pct) * color.g / normalize + pct,
		       (normalize - pct) * color.b / normalize }

	gl_FragColor = adjustedRGB(color)
}
