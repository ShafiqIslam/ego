package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func method() {
	fmt.Println("methods are just functions, with a receiver parameter")

	v := Vertex{3, 4}
	fmt.Println(v.Abs())

}

func methodOnNonStructs() {
	fmt.Println("methods can be on non struct type also")
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func methodOnPointer() {
	fmt.Println("methods often need to modify their receiver, pointer are more common than value")

	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

func main() {
	method()
	fmt.Println()
	methodOnNonStructs()
	fmt.Println()
	methodOnPointer()
}
