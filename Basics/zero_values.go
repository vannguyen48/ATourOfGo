// Variables declared without an explicit value are given their zero values.
// 0, false, "" are examples of zero values.
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
