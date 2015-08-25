# bootgo

A barebones OS kernel written in go

The kernel contains some code modified from devos barebones tutorial http://wiki.osdev.org/Bare_Bones

## setup
To compile bootgo You need a gccgo cross-compiler(my gccgo version is 5.2.0)

1. build a target i386/i686 gcc cross-compiler with go enabled, follow the article http://wiki.osdev.org/GCC_Cross-Compiler (gcc 5.2.0 is recommended)

2. install nasm from your repositories

3. install qemu for test

## compiler & run!

1. compile: `make GCC=i686-elf-gcc GCCGO=i686-elf-gccgo`, replace `GCC` and `GCCGO` with your binary name

2. run on qemu: `make run-qemu QEMU=qemu-system-i386`, replace `QEMU` with your target binary name
