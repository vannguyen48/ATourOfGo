// In Go, a name is exported if it begins with a capital letter. 
// For example, "pizza" and "pi" are not exported because they do not start with a capital letter.
// Resource: math library constants:  https://golang.org/pkg/math/#pkg-constants

package main

import (
	"fmt"
	"math"
)

func main() {
	// This won't work because "pi" is not a constant.
	// fmt.Println(math.pi)

	// This will work because "Pi" is a constant.
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt2)
	fmt.Println(math.Ln2)
	fmt.Println(math.Ln10)
}
