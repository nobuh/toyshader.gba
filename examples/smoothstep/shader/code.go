package shader

func plot(st vec2, pct int16) int16 {
	return smoothstep(pct - 5, pct, st.y) -
	       smoothstep(pct, pct + 5, st.y)
}

func shader() {
	st := normalizedXY(gl_FragCoord)

	// gradient
	y := smoothstep(normalize * 2/10, normalize * 8/10, st.x)
	color := vec3{ y, y, y }

	// line
	pct := plot(st, y)
	color = vec3 { (normalize - pct) * color.r / normalize,
		       (normalize - pct) * color.g / normalize + pct,
		       (normalize - pct) * color.b / normalize }

	gl_FragColor = adjustedRGB(color)
}
