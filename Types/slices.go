// A slice is dynamically sized. Slices are much nore common than array
// A slice is formed by specifying two indices, a low and high bound, separated by a colon
// Slices are like references to arrays. A slice describes a section of an underlying array
// Changing the elements of a slice modifies the corresponding elements of its underlying array.
// When slicing, you may omit the high or low bounds to use their defaults instead.
// The default is zero for the low bound and the length of the slice for the high bound.
// The zero value of a slice is nil. A nil slice has a length of 0 and has no underlying array.

package main

import (
	"fmt"
	"strings"
)

func createSliceWithMake() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func tictactoe() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendSlice() {
	var s []int
	s = append(s, 0)
	fmt.Println(s)
	s = append(s, 1)
	fmt.Println(s)
	s = append(s, 2, 3, 4)
	fmt.Println(s)
}

func rangeDemo() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Includes the first element, excludes the last one
	// This prints out 3 5 7
	var s []int = primes[1:4]
	fmt.Println(s)
	favoriteWords := [4]string{"Apple", "Bees", "Coconut", "Dairy"}
	fmt.Println(favoriteWords)
	favoriteWords[3] = "Dogs"
	fmt.Println(favoriteWords)
	// A slice literal is like an array literal without the length.
	// The two below are identical, one is an array literal
	// the other creates the same array as above, then builds a slice that references it.
	array1 := [3]bool{true, true, false}
	slice1 := []bool{true, true, false}
	fmt.Println(array1)
	fmt.Println(slice1)
	m := []int{2, 3, 5, 7, 11, 13}

	m1 := m[1:4]
	fmt.Println(m1)

	m2 := m[:2]
	fmt.Println(m2)

	m3 := m[1:]
	fmt.Println(m3)
	// The length of a slice is the number of elements it contains
	// The capacity of a slice is the number of elements in the underlying aray, counting from the first element in the slice

	fmt.Printf("len=%d cap=%d %v\n", len(m), cap(m), m)

	// Create dynamically sliced array with the built-in make function.
	// The make dunction allocates a zeroed array and returns a slice that refers to that array.
	createSliceWithMake()

	tictactoe()

	appendSlice()

	rangeDemo()
}
