package shader

func shader() {
	p := normalizedXY(gl_FragCoord)

	const n = int16(normalize)

	left := step(3, p.x)
	bottom := step(3, p.y)
	center := step(42, p.x) + step(42, n - p.x)
	center2 := step(42, p.y) + step(42, n - p.y)
	top := step(3, n - p.y)
	right := step(3, n - p.x)

	// left * bottom is similar to logical and 
	// also * top * right
	color := vec3{ left * bottom /n * top /n * right /n * center /n * center2 /n,
		       left * bottom /n * top /n * right /n * center /n * center2 /n,
		       left * bottom /n * top /n * right /n * center /n * center2 /n }

	gl_FragColor = adjustedRGB(color)
}
