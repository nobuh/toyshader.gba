package shader

func shader() {
	// if 255 * 240, the value is over int16, so divide 255 to 51 * 5 
	st := vec2{ cdepth / u_resolution.x * gl_FragCoord.x,
		    cdepth / u_resolution.y * gl_FragCoord.y }
	gl_FragColor = vec3{ st.x, st.y, 0}
}
