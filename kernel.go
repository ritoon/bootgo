package kernel

import "unsafe"

const (
  COLOR_BLACK = iota
	COLOR_BLUE
	COLOR_GREEN
	COLOR_CYAN
	COLOR_RED
	COLOR_MAGENTA
	COLOR_BROWN
	COLOR_LIGHT_GREY
	COLOR_DARK_GREY
	COLOR_LIGHT_BLUE
	COLOR_LIGHT_GREEN
	COLOR_LIGHT_CYAN
	COLOR_LIGHT_RED
	COLOR_LIGHT_MAGENTA
	COLOR_LIGHT_BROWN
	COLOR_WHITE
)

const (
  VGA_WIDTH = 80
  VGA_HEIGHT = 25
)

var row, column, color uint8
var buffer uintptr

func makeColor(fg uint8, bg uint8) uint8 {
  return fg | bg << 4
}

func makeVGAEntry(c byte, color uint8) uint16 {
  return uint16(c) | uint16(color) << 8
}

func terminalInit() {
  row = 0
  column = 0
  color = makeColor(COLOR_LIGHT_GREY, COLOR_BLACK)
  buffer = 0xB8000
  for y := 0; y < VGA_HEIGHT; y += 1 {
    for x := 0; x < VGA_WIDTH; x += 1 {
      terminalPutEntryAt(' ', color, uint8(x), uint8(y))
    }
  }
}

func terminalSetColor(c uint8) {
  color = c
}

func terminalPutEntryAt(c byte, color uint8, x uint8, y uint8){
  index := y * VGA_WIDTH + x
  addr := (*uint16)(unsafe.Pointer(buffer + 2 * uintptr(index)))
  *addr = makeVGAEntry(c, color)
}

func terminalPutChar(c byte) {
  terminalPutEntryAt(c, color, column, row)
  column+=1
  if column == VGA_WIDTH {
    column = 0
    row += 1
    if row == VGA_HEIGHT {
      row = 0
    }
  }
}

func writeString(data string) {
  for i:= 0; i < len(data); i+=1 {
    terminalPutChar(data[i])
  }
}

func Main() {
  terminalInit()
  writeString("hello, kernel!")
}
