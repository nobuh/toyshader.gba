toyshader: toyshader.bin
	VisualBoyAdvance -1 toyshader.bin
toyshader.bin: main.go shader/engine.go shader/code.go
	tinygo build -target gameboy-advance -o toyshader.bin main.go
clean:
	rm *.bin 
