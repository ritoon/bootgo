ENTRY(_start)

SECTIONS
{
  . = 1M;

  .multiboot :
  {
    *(.multiboot)
  }

  .text :
  {
    *(.text)
  }

  .data :
  {
    *(.data)
  }
  
  .bss :
  {
    *(.boot_stack)
  }
}
