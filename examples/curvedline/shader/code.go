package shader

func plot(st vec2, pct int16) int16 {
	return smoothstep(pct - 4, pct, st.y) -
	       smoothstep(pct, pct + 4, st.y)
}

func shader() {
	st := normalizedXY(gl_FragCoord)

	// gradient
	y := pow(st.x, 2)
	color := vec3{ y, y, y }

	// line
	pct := plot(st, y)
	color = vec3 { (normalize - pct) * color.r / normalize,
		       (normalize - pct) * color.g / normalize + pct,
		       (normalize - pct) * color.b / normalize }

	gl_FragColor = adjustedRGB(color)
}
