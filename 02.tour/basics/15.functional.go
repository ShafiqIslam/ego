package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func firstClassFunctions() {
	fmt.Println("functions as function arguments and return values")

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closures() {
	fmt.Println("function may access and assign to the referenced variables outside its body")

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		defer func() {
			r := a + b
			a = b
			b = r
		}()
		return a
	}
}

func functionalExercise() {
	fmt.Println("an exercise of functional program with fibonacci")

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func main() {
	firstClassFunctions()
	fmt.Println()
	closures()
	fmt.Println()
	functionalExercise()
}
