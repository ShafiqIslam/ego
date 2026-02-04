package main

import (
	"fmt"
	"math"
	"math/rand"
)

func factored() {
	fmt.Println("groups the imports into a parenthesized, \"factored\" import statement")
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func importing() {
	fmt.Println("can have multiple import statements")
	fmt.Println("My favorite number is", rand.Intn(10))
}

func names() {
	fmt.Println("a name is exported if it begins with a capital letter")
	fmt.Println(math.Pi)
	// fmt.Println(math.pi)
}

func main() {
	importing()
	fmt.Println()
	factored()
	fmt.Println()
	names()
}
