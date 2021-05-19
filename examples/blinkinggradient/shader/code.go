package shader

func shader() {
	// if 255 * 240, the value over int16, so divide 255 to 51 * 5 
	st := vec2{ abs(cos(u_time * 10)) * 51/100 * gl_FragCoord.x / resolution_x * 5,
		    abs(sin(u_time * 10)) * 51/100 * gl_FragCoord.y / resolution_y * 5 }
	gl_FragColor = vec4{ 0, st.x, st.y, 255}
}
