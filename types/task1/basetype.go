package main

import (
	"fmt"
)

func main() {
	var i8 int8 = -128
	var u8 uint8 = 255
	var i16 int16 = -32768
	var u16 uint16 = 65535
	var i32 int32 = -2147483648
	var u32 uint32 = 4294967295
	var i64 int64 = -9223372036854775808
	var u64 uint64 = 18446744073709551615
	var r rune = 'A'

	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	var c64 complex64 = complex(1, 2)
	var c128 complex128 = complex(3, 4)

	var b byte = 'B'
	var s string = "Базовые типы данных"

	var bo bool = true

	fmt.Printf("int8: %d\n", i8)
	fmt.Printf("uint8: %d\n", u8)
	fmt.Printf("int16: %d\n", i16)
	fmt.Printf("uint16: %d\n", u16)
	fmt.Printf("int32: %d\n", i32)
	fmt.Printf("uint32: %d\n", u32)
	fmt.Printf("int64: %d\n", i64)
	fmt.Printf("uint64: %d\n", u64)
	fmt.Printf("rune: %c\n", r)
	fmt.Printf("float32: %f\n", f32)
	fmt.Printf("float64: %.15f\n", f64)
	fmt.Printf("complex64: %v\n", c64)
	fmt.Printf("complex128: %v\n", c128)
	fmt.Printf("byte: %c\n", b)
	fmt.Printf("string: %s\n", s)
	fmt.Printf("bool: %t\n", bo)
}
