GO111MODULE=off

toyshader: toyshader.bin
	mgba -1 toyshader.bin
toyshader.bin: main.go shader/engine.go shader/code.go
	# opt s seems to be a litte slow at tinygo 0.21 and opt z is very slow and unexpected behavior
	tinygo build -opt 2 -target gameboy-advance -o toyshader.bin main.go
clean:
	rm *.bin *.sav
