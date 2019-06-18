// A map maps keys to values.
// Zero value of map is nil. a nil map has no keys and keys cannot be added
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func mapLiterals() {
	// map literals are like struct literals, but the keys are required.
	var m1 = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println("Map literals: ", m1)
	
	// If the top level is a type name, you can omit it from the elements of the literal.
	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)
}

// For the most part, inserting/updating/accessing an element in a map is very similar to
// doing those operations with a dictionary.
func mutatingMaps() {
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// Deleting an element.
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// if key is in m, ok is true. If not, ok is false.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	mapLiterals()
	mutatingMaps()
}
