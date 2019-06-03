// Go has only one looping construct, the for loop.
// The for loop has similar components to the for loop in any given language.
// No parentheses surrounding the for loop.
/*
This below is an example of an infinite loop
for {
}
*/
package main

import "fmt"

func forloop() {
	// The init and post statements are optional.
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func forloopaswhile() {
	// For is also Go's while statement.
	sum := 1
	for sum < 500 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	sum := 0;
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
	forloop()
	forloopaswhile()
}