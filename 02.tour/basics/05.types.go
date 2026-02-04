package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func basicTypes() {
	fmt.Println("Accessing type of value")

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zero() {
	fmt.Println("Variables declared without explicit initial value are given their zero value.")

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func conversion() {
	fmt.Println("Assignment between different type requires an explicit conversion")

	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}

func inference() {
	fmt.Println("Types are infered when assigned")

	var i = 42        // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	fmt.Printf("i is of type %T\n", i)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}

func main() {
	basicTypes()
	fmt.Println()
	zero()
	fmt.Println()
	conversion()
	fmt.Println()
	inference()
}
