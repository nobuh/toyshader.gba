rgb: rgb.bin
	VisualBoyAdvance -1 rgb.bin
crt.o: crt.S
	arm-linux-gnueabi-gcc -c crt.S
rgb.o: rgb.c
	arm-linux-gnueabi-gcc -Wall -c rgb.c -o rgb.o
rgb.out: rgb.o crt.o cart.ls
	arm-linux-gnueabi-ld -o rgb.out -T cart.ls crt.o rgb.o
rgb.bin: rgb.out
	arm-linux-gnueabi-objcopy -O binary rgb.out rgb.bin
clean:
	rm *.bin *.o *.out
