// Packages: Every go program is made up of packages and starts running in package main.
// By convention, the package name is the same as the last element of the import path.
// For example, the "math/rand" package comprises files that begin with the statement package rand.

// External resources: https://golang.org/pkg/math/rand/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// The environment in which these programs are executed is deterministic, this will return the same number every time.
	fmt.Println("Selected number is ", rand.Intn(10))

	// This will generate different number every time.
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Selected number after seeding is ", rand.Intn(10))
}