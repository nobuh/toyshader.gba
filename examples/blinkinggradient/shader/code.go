package shader

func shader() {
	// cdepth = 51 * 5
	st := vec2{ abs(cos(u_time * 10)) * 51 / 100 * gl_FragCoord.x / u_resolution.x * 5,
		    abs(sin(u_time * 10)) * 51 / 100 * gl_FragCoord.y / u_resolution.y * 5 }
	gl_FragColor = vec3{ 0, st.x, st.y}
}
