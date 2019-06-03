// A function can take 0 or more arguments. notice that the type comes after the variable names.

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// When two or more consecutuve named function parameters share a type, you can omit the type from all but the last.
func add1(x, y int) int {
	return x + y
}

func subtract (x int, y int) int {
	return x - y
}
// A function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}
// Go's return values may be named, they are at the top of the function.
// These names shoyld be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as 'naked' return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(subtract(200, 80))
	fmt.Println(add(116, 100))
	fmt.Println(add1(116, 100))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(52))
}
