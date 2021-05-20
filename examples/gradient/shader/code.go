package shader

func shader() {
	// if 255 * 240, the value is over int16, so divide 255 to 51 * 5 
	st := vec2{ int16(cdepth * uint16(gl_FragCoord.x) / resolution_x),
		    int16(cdepth * uint16(gl_FragCoord.y) / resolution_y) }
	gl_FragColor = vec4{ st.x, st.y, 0, 255}
}
