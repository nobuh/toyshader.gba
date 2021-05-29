package shader

/* https://thebookofshaders.com/06/
vec3 colorA = vec3(0.149,0.141,0.912);
vec3 colorB = vec3(1.000,0.833,0.224);

void main() {
    vec3 color = vec3(0.0);

    float pct = abs(sin(u_time));

    // Mix uses pct (a value from 0-1) to
    // mix the two colors
    color = mix(colorA, colorB, pct);

    gl_FragColor = vec4(color,1.0);
}
*/

var colorA vec3 = vec3 { 14, 14, 70 }
var colorB vec3 = vec3 { 80, 65, 20 }

func shader() {
	ratio := abs(sin(u_time * 10))

	color := mix(colorA, colorB, ratio)

	gl_FragColor = adjustedRGB(color)
}
