func rand() int16 {
        // based on https://gist.github.com/johansten/3633917
        a := fract(dot(gl_FragCoord, vec2{ 206, 1245 })) - resolution_y /2
        s := a * (618 + a * a * (-3802 + a * a * 5339))
        t := fract(s * 4735)
        return t
}

func pseudoShader() {
	c := rand() * 255 / 100
	gl_FragColor = vec4{c, c, c, 255}
}
