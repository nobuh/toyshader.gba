toyshader: toyshader.bin
	VisualBoyAdvance toyshader.bin
toyshader.bin: toyshader.go
	tinygo build -target gameboy-advance -o toyshader.bin toyshader.go
clean:
	rm *.bin 
