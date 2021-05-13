// toyshader.c

typedef unsigned short hword;	// half word

#define VRAM		0x06000000		// VRAM frame buffer start
#define gba_register(p)	*((volatile hword*) p)	// I/O register volatile handling
#define LCD_CTRL	0x04000000		// LCD control
#define	LCD_BG2EN	0x0400			// Enable BG2
#define	LCD_MODE3	3			// Video Mode 3
#define	LCD_WIDTH	240
#define LCD_HEIGHT	160
#define	CMAX		31			// Max Color Value

typedef struct {
        hword r;
        hword g;
        hword b;
        hword a;
} Vec4;

Vec4 ivec4(hword r, hword g, hword b, hword a) {
	Vec4 v;
	v.r = r;
	v.g = g;
	v.b = b;
	v.a = a;
	return v;
}

Vec4 gl_FragColor = { 0, 0, 0, 0 };  


void PseudoShader() {
	// 240x160 screen and color values in 0..31

	gl_FragColor = ivec4(CMAX/2, CMAX/2, CMAX/2, 1);

}


int main() {
	hword* ptr;
	int i;
        hword r;

	// init LCD

	gba_register(LCD_CTRL) = LCD_BG2EN | LCD_MODE3;

	// Draw Loop
	while (1) {
		ptr = (hword*) VRAM;

		for (i = 0; i < LCD_HEIGHT * LCD_WIDTH; i++) {

        		PseudoShader();

        		r = gl_FragColor.r;
        		r += (gl_FragColor.g << 5);
        		r += (gl_FragColor.b << 10);

			*(ptr++) = r;
		}
	}
}
