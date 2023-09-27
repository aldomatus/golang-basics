// In Go (Golang), the language offers a range of primitive data types that
// serve as the building blocks for creating complex data structures and
// variables. Below is a summary of the most common primitive data types:

// Integer Types:
// uint: Unsigned integer; holds only positive integers. Size depends on machine
// architecture (32 or 64 bits).
// uint8: An 8-bit unsigned integer (range: 0 to 255).
// uint16: A 16-bit unsigned integer (range: 0 to 65,535).
// uint32: A 32-bit unsigned integer (range: 0 to 4,294,967,295).
// uint64: A 64-bit unsigned integer (range: 0 to 18,446,744,073,709,551,615).

// Floating-Point Types:
// float32: A 32-bit floating-point number.
// float64: A 64-bit floating-point number.

// String Types:
// string: Immutable sequence of characters. Unlike in some languages like
// Python, Go strings are not lists of characters and are immutable by default.

// Complex Numbers:
// complex64: A complex number with float32 real and imaginary parts.
// complex128: A complex number with float64 real and imaginary parts.

// Other Numeric Types:
// int: Signed integers; size is architecture-dependent (32 or 64 bits).
// int8, int16, int32, int64: Similar to unsigned integers but can hold negatives.
// byte: Alias for uint8.
// rune: Alias for int32, represents a Unicode character.

// Go also has a boolean type (bool) for true/false values and derived types like
// arrays, slices, and maps, but these are not considered primitive types.

package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	// Integer Types
	var myUint uint = 10
	var myUint8 uint8 = 255
	var myUint16 uint16 = 65535
	var myUint32 uint32 = 4294967295
	var myUint64 uint64 = 18446744073709551615

	// Floating-Point Types
	var myFloat32 float32 = 3.14
	var myFloat64 float64 = 3.141592653589793

	// String Type
	var myString string = "Hello, Go!"

	// Complex Numbers
	var myComplex64 complex64 = 1 + 2i
	var myComplex128 complex128 = cmplx.Sqrt(-5 + 12i)

	// Other Numeric Types
	var myInt int = -100
	var myInt8 int8 = -128
	var myInt16 int16 = -32768
	var myInt32 int32 = -2147483648
	var myInt64 int64 = -9223372036854775808

	// Byte and Rune
	var myByte byte = 'A' // alias for uint8
	var myRune rune = 'あ' // alias for int32

	// Output each variable
	fmt.Println("myUint:", myUint)
	fmt.Println("myUint8:", myUint8)
	fmt.Println("myUint16:", myUint16)
	fmt.Println("myUint32:", myUint32)
	fmt.Println("myUint64:", myUint64)
	fmt.Println("myFloat32:", myFloat32)
	fmt.Println("myFloat64:", myFloat64)
	fmt.Println("myString:", myString)
	fmt.Println("myComplex64:", myComplex64)
	fmt.Println("myComplex128:", myComplex128)
	fmt.Println("myInt:", myInt)
	fmt.Println("myInt8:", myInt8)
	fmt.Println("myInt16:", myInt16)
	fmt.Println("myInt32:", myInt32)
	fmt.Println("myInt64:", myInt64)
	fmt.Println("myByte:", myByte)
	fmt.Println("myRune:", myRune)
}
