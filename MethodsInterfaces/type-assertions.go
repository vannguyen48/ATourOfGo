// A type assertion provides access to an interdace value's 
// underlying concrete value.
// A type switch is a construct that permits several type of assertions
// in series. Like a regular switch statement, but the case in a type switch
// specify type.
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	var i interface{} = "hello"
	// value i holds  the concrete type string
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)
	// value i holds the concrete type float
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// f = i.(float64) // panic
	fmt.Println(f)

	do(21)
	do("hello")
	do(true)
}
