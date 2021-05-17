func pseudoShader() {
	// if 255 * 240, the value is over int16, so divide 255 to 51 * 5 
	st := vec2{ 51 * gl_FragCoord.x / resolution_x * 5,
		    51 * gl_FragCoord.y / resolution_y * 5 }
	gl_FragColor = vec4{ st.x, st.y, 0, 255}
}
