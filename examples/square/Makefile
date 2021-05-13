square: square.bin
	VisualBoyAdvance -1 square.bin
square.bin: square.go
	tinygo build -target gameboy-advance -o square.bin square.go
