package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i8 int8
	var u8 uint8
	var i16 int16
	var u16 uint16
	var i32 int32
	var u32 uint32
	var i64 int64
	var u64 uint64
	var r rune
	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128
	var b byte
	var s string
	var bo bool

	fmt.Printf("int8: %d bytes\n", unsafe.Sizeof(i8))
	fmt.Printf("uint8: %d bytes\n", unsafe.Sizeof(u8))
	fmt.Printf("int16: %d bytes\n", unsafe.Sizeof(i16))
	fmt.Printf("uint16: %d bytes\n", unsafe.Sizeof(u16))
	fmt.Printf("int32: %d bytes\n", unsafe.Sizeof(i32))
	fmt.Printf("uint32: %d bytes\n", unsafe.Sizeof(u32))
	fmt.Printf("int64: %d bytes\n", unsafe.Sizeof(i64))
	fmt.Printf("uint64: %d bytes\n", unsafe.Sizeof(u64))
	fmt.Printf("rune: %d bytes\n", unsafe.Sizeof(r))
	fmt.Printf("float32: %d bytes\n", unsafe.Sizeof(f32))
	fmt.Printf("float64: %d bytes\n", unsafe.Sizeof(f64))
	fmt.Printf("complex64: %d bytes\n", unsafe.Sizeof(c64))
	fmt.Printf("complex128: %d bytes\n", unsafe.Sizeof(c128))
	fmt.Printf("byte: %d bytes\n", unsafe.Sizeof(b))
	fmt.Printf("string: %d bytes\n", unsafe.Sizeof(s))
	fmt.Printf("bool: %d bytes\n", unsafe.Sizeof(bo))
}
