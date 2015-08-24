section .data
  MAGIC equ 0x1BADB002
  FLAGS equ 0
  CHECKSUM equ -(MAGIC + FLAGS)

section .multiboot
align 4
  dd MAGIC
  dd FLAGS
  dd CHECKSUM

section .boot_stack
align 4
stack_bottom:
resb 16384
stack_top:

section .text
global _start
_start:
  extern bootgo.kernel.Main
  call bootgo.kernel.Main
  cli

.hang:
  hlt
  jmp .hang
