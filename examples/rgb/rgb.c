typedef unsigned short hword;	// half word

#define VRAM		0x06000000		// VRAM frame buffer start
#define gba_register(p)	*((volatile hword*) p)	// I/O register volatile handling
#define LCD_CTRL	0x04000000		// LCD control
#define	LCD_BG2EN	0x0400			// Enable BG2
#define	LCD_MODE3	3			// Video Mode 3
#define	LCD_WIDTH	240
#define LCD_HEIGHT	160
#define	BGR(r, g, b)	((b << 10) + (g << 5) + r)

int main() {
	hword* ptr = (hword*) VRAM;
	int i, j;

	// init LCD

	gba_register(LCD_CTRL) = LCD_BG2EN | LCD_MODE3;

	// Tricolor
		
	for (i = 0; i < (LCD_HEIGHT / 3); i++) {
		for (j = 0; j < LCD_WIDTH; j++) {
			*(ptr++) = BGR(31, 0, 0);
		}
	}
	for (i = 0; i < (LCD_HEIGHT / 3); i++) {
		for (j = 0; j < LCD_WIDTH; j++) {
			*(ptr++) = BGR(0, 31, 0);
		}
	}
	for (i = 0; i < (LCD_HEIGHT / 3); i++) {
		for (j = 0; j < LCD_WIDTH; j++) {
			*(ptr++) = BGR(0, 0, 31);
		}
	}

	// forever
	while (1) ;
}
