# Basic Linux Process

## Start
Compile the simple program `main.c` without assembling the output - e.g., by running `make`.
You will get a variation of the assembly code in the next section, depending on what processor you're running on.
In my case, I'm running on a `Raspberry Pi 5 Model B Rev 1.0` machine with an `armv8-a` architechture.

## Assembly Code
```armasm
	.arch armv8-a
	.file	"main.c"
	.text
	.global	A
	.data
	.align	2
	.type	A, %object
	.size	A, 4

A:
	.word	356
	.section	.rodata
	.align	3
.LC0:
	.string	"'c' is %d"
	.text
	.align	2
	.global	main
	.type	main, %function

main:
.LFB6:
	.cfi_startproc
  ; armv8-a provides 31 general purpose registers and supports access in 2 forms: 32-bit & 64-bit
  ; `wNUM` signifies 32-bit access & `xNUM` 64-bit access
  ; x30 is used as the link register: stores the address to return to when a subroutine call completes
  stp	x29, x30, [sp, -32]!
	.cfi_def_cfa_offset 32
	.cfi_offset 29, -32
	.cfi_offset 30, -24
	mov	x29, sp
	mov	w0, 105
	str	w0, [sp, 20]
	mov	x0, 4
	bl	malloc
	str	x0, [sp, 24]
	adrp	x0, A
	add	x0, x0, :lo12:A
	ldr	w1, [x0]
	ldr	w0, [sp, 20]
	add	w1, w1, w0
	ldr	x0, [sp, 24]
	str	w1, [x0]
	ldr	x0, [sp, 24]
	ldr	w0, [x0]
	mov	w1, w0
	adrp	x0, .LC0
	add	x0, x0, :lo12:.LC0
	bl	printf
	ldr	x0, [sp, 24]
	bl	free
	mov	w0, 0
	ldp	x29, x30, [sp], 32
	.cfi_restore 30
	.cfi_restore 29
	.cfi_def_cfa_offset 0
	ret
	.cfi_endproc
.LFE6:
	.size	main, .-main
	.ident	"GCC: (Ubuntu 13.3.0-6ubuntu2~24.04) 13.3.0"
	.section	.note.GNU-stack,"",@progbits

```


## References
- [Armv8-A Instruction Set Architecture](https://developer.arm.com/-/media/Arm%20Developer%20Community/PDF/Learn%20the%20Architecture/Armv8-A%20Instruction%20Set%20Architecture.pdf)
