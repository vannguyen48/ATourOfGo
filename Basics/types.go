// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32
     // represents a Unicode code point

// float32 float64

// complex64 complex128

/*
By default, variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric type
false for boolean type
"" the empty string for strings

The expression T(v) converts v to the T type
*/

// When declaring a variable without specifying an explicit type, the variable's type is inferred
// from the value on the right hand side.
package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var zz uint = uint(f)
	fmt.Println(x, y, zz)

	// Type conversion.
	var i int = 23
	var fl float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i, fl, u)

	// Type inference
	v := 42.5236425 // change me!
	fmt.Printf("v is of type %T\n", v)
}
