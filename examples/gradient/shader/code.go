package shader

func shader() {
	st := vec2{ cdepth / u_resolution.x * gl_FragCoord.x,
		    cdepth / u_resolution.y * gl_FragCoord.y }
	gl_FragColor = vec3{ st.x, st.y, 0}
}
