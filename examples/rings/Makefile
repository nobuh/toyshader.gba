toyshader: toyshader.bin
	VisualBoyAdvance -1 toyshader.bin
toyshader.bin: toyshader.go
	tinygo build -target gameboy-advance -o toyshader.bin toyshader.go
clean:
	rm *.bin 
