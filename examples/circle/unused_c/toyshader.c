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

typedef struct {
	hword x;
	hword y;
} Vec2;

Vec2 ivec2(hword x, hword y) {
	Vec2 v;
	v.x = x;
	v.y = y;
	return v;
}

Vec2 gl_FragCoord = { 0, 0 };

hword sqrt(hword x){
	hword nMAX = 5;

    	if (x <= 0) return 0;

	hword t = x;
 
    	for(hword n=1; n<=nMAX; n++){
        	t = (t + x / t) / 2;
    	}
 
    	return (hword)t;
}

hword length(Vec2 a, Vec2 b) {
        return sqrt( (a.x - b.x)*(a.x - b.x) + (a.y - b.y)*(a.y - b.y) );
}

/* Your Shader Code is Here */

void PseudoShader() {
	// LCD_WIDTH  240
	// LCD_HEIGHT 160
	// CMAX       Max Value of RGB

        Vec2 pos = ivec2(
                gl_FragCoord.x - LCD_WIDTH/2, 
                gl_FragCoord.y - LCD_HEIGHT/2
        );

        Vec2 o = { LCD_WIDTH/2, LCD_HEIGHT/2 };

        hword r = 10;
        hword len = length(pos, o);
	hword c = r / len;

        gl_FragColor = ivec4(c, c, c, 1) ;

}

/* Until This */

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
