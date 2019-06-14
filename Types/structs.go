// A struct is a collection of fields
// Struct fields are accessed using a dot.
// Struct fields can be accessed through a struct pointer
// To access the field X of a struct when we have the struct pointer p we could write (*p).X, or just p.X
// Struct literal denotes a newly allocated struct value by listing the values of its fields.
// You can list a subset of fields by using the Name: syntax

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

type Coordinates struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p1 = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Coordinates{5, 7})
	v := Vertex{5, 9}
	v.X = 13
	fmt.Println(v.X)
	p := &v
	p.X = 103
	fmt.Println(p.X)
	fmt.Println(v1, p1, v2, v3)
}
