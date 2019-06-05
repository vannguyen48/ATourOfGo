// Flow-Control section, #8

// Go exercise: Given a number x, find the number z for which z^2 is most nearly x.
// Computers compute the square root using a loop, so we'll leverage that in this exercise.
// Use the math.Sqrt function in the standard library to see how close we are to the library solution.

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
	for i := 1; i < int(x/2); i++ {
		current_z := z
		z -= (z*z - x) / (2*z)
		if  current_z - z == 0 {
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(27), math.Sqrt(27)) // 5.196152422706632 5.196152422706632
	fmt.Println(Sqrt(25), math.Sqrt(25)) // 5 5
	fmt.Println(Sqrt(100), math.Sqrt(100)) // 10 10
	fmt.Println(Sqrt(300), math.Sqrt(300)) // 17.320508075688775 17.320508075688775
}
