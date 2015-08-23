package kernel

import "unsafe"

//type VGAColor uint8

const (
  COLOR_BLACK = 0
	COLOR_BLUE = 1
	COLOR_GREEN = 2
	COLOR_CYAN = 3
	COLOR_RED = 4
	COLOR_MAGENTA = 5
	COLOR_BROWN = 6
	COLOR_LIGHT_GREY = 7
	COLOR_DARK_GREY = 8
	COLOR_LIGHT_BLUE = 9
	COLOR_LIGHT_GREEN = 10
	COLOR_LIGHT_CYAN = 11
	COLOR_LIGHT_RED = 12
	COLOR_LIGHT_MAGENTA = 13
	COLOR_LIGHT_BROWN = 14
	COLOR_WHITE = 15
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
      index := y * VGA_WIDTH + x;
      addr := (*uint16)(unsafe.Pointer(buffer + uintptr(2 * index)))
      *addr = makeVGAEntry(' ', color)
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
  writeString("hello! kernel in go!!")
}
