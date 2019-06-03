// If statements are not surrounded by () but the braces are required.

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
// Complex number.
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// If with a short statement.
// Variables declared by the statement are only in scope in the end of if.
// That means you can't return v where you're return lim
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(sqrt(9), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
