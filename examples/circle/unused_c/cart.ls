OUTPUT_ARCH(arm)

SECTIONS {
	.text 0x08000000 : { *(.text) }
	.data	: { *(.data) }
	.rodata	: { *(.rodata*) }
	.bss	: { *(.bss) }
	}
