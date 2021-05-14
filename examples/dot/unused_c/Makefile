toyshader: toyshader.bin
	VisualBoyAdvance -1 toyshader.bin
crt.o: crt.S
	arm-linux-gnueabihf-gcc -c crt.S
toyshader.o: toyshader.c
	#arm-linux-gnueabi-gcc -Wall -O2 -nostdlib -nostartfiles -ffreestanding -c toyshader.c -o toyshader.o 
	arm-linux-gnueabihf-gcc -Wall -O2 -ffreestanding -c toyshader.c -o toyshader.o 
toyshader.out: toyshader.o crt.o cart.ls
	arm-linux-gnueabihf-ld -o toyshader.out -T cart.ls crt.o toyshader.o
toyshader.bin: toyshader.out
	arm-linux-gnueabihf-objcopy -O binary toyshader.out toyshader.bin
clean:
	rm *.bin *.o *.out
