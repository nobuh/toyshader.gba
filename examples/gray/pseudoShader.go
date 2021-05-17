func pseudoShader() {
	gray := int16(255/2)
	half_transparent := int16(255/2)
	gl_FragColor = vec4{gray, gray, gray, half_transparent}
}
