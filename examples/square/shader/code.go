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
	sum := left * bottom /n * top /n * right /n * center /n * center2 /n

	var color vec3
	if p.x < 42 && p.y > 42 {
		color = vec3{ 0, 0, sum }
	} else if p.x > 42 && p.y < 42 {
		color = vec3{ sum, 0, 0 }
	} else {
		color = vec3{ sum, sum, sum }
	}

	gl_FragColor = adjustedRGB(color)
}
