// RGB15 : GBA's Blue Green Red 15 bit format shader
//
// 0x0BBBBBGGGGGRRRRR : 15 bit BGR format in little endian

#include <stdio.h>
#include <stdlib.h>
#include <math.h>

#define MAXCOLOR	31 
int pixel_idx;

typedef struct {
	float x;
	float y;
} Vec2;

Vec2 vec2(float x, float y) {
	Vec2 v = { 0, 0 };
	v.x = x;
	v.y = y;
	return v;
}

Vec2 resolution = { 240, 160 };

Vec2 gl_FragCoord = { 0, 0 }; 

typedef struct {
	float r;
	float g;
	float b;
	float a;
} Vec4;

Vec4 gl_FragColor = { 0, 0, 0, 1 }; 

Vec4 vec4(float r, float g, float b, float a) {
	Vec4 v = { 0, 0, 0, 0 }; 
	v.r = r;
	v.g = g;
	v.b = b;
	v.a = a;
	return v;
}	

float length(Vec2 a, Vec2 b) {
	return sqrtf( (a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y) );
} 	

void PseudoShader() {

 	const float ncircle = 20.0;
 	const float orbit_radius = 0.4;
 	const float circle_radius = 0.01;

	Vec2 pos = vec2(
		(gl_FragCoord.x - 0.5 * resolution.x) / resolution.y, 
		(gl_FragCoord.y - 0.5 * resolution.y) / resolution.y 
	);

	Vec4 color = { 0, 0, 0, 1 };

 	for(float angle = 0.0; angle < 2.0 * 3.14; angle += 2.0 * 3.14 / ncircle) {
         	float x = orbit_radius * cos(angle);
         	float y = orbit_radius * sin(angle);
		float len = length(pos, vec2(x, y));

		if (len < circle_radius) len = circle_radius;

         	color.r += circle_radius / len;
         	color.g += circle_radius / len;
         	color.b += circle_radius / len;
	}
	gl_FragColor = color;
}

int RGB15() {
	int b, g, r, rgb15;

	PseudoShader();

	b = (int)abs(MAXCOLOR * gl_FragColor.b * gl_FragColor.a);
	g = (int)abs(MAXCOLOR * gl_FragColor.g * gl_FragColor.a);
	r = (int)abs(MAXCOLOR * gl_FragColor.r * gl_FragColor.a);

	rgb15 = r;		// 5 bit RED value is 0..31
	rgb15 += (g << 5);	// add Green
	rgb15 += (b << 10); 	// add Blue

	return rgb15;
}

int main(int argc, char** argv) {
	FILE *dst;
	int i, buf;
	
	dst = fopen(argv[1], "w");
	if (dst == NULL) {
		printf("Destination file open failure\n");
		return 1;
	}

	for (pixel_idx = 0; pixel_idx < resolution.x * resolution.y; pixel_idx++) {

		gl_FragCoord.x = (float)(pixel_idx % (int)resolution.x);
		gl_FragCoord.y = (float)(pixel_idx / (int)resolution.x);

		buf = RGB15();
		
		// write 16 bit value in little endian
		if (fputc(buf & 0xFF, dst) == EOF) {
			printf("File write failure\n");
			goto error;
		}
		if (fputc(buf >> 8, dst) == EOF) {
			printf("File write failure\n");
			goto error;
		}	
	}

	fclose(dst);
	return 0;

error:
	fclose(dst);
	return 1;
}
