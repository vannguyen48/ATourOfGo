// The var statement declares a list of variables, as ub function argument lists, the type is last.
// The var statement can be at package or function level.

package main

import "fmt"

var c, python, java bool
// Variables with initializers
var k, j int = 1, 2
// Short variable declarations - this mist be inside a function
// short := "short varable"

func main() {
	var i int
	short := "short varable"
	fmt.Println(i, c, python, java)
	fmt.Println(k, j)
	fmt.Println(short)
}
