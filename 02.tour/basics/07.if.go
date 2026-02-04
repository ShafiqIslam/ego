package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func basicIf() {
	fmt.Println("Find sqrt with basic if")

	fmt.Println(sqrt(2), sqrt(-4))
}

func ifWithInitializer() {
	fmt.Println("Find pow with initializer variable")
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

func powV2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func ifElse() {
	fmt.Println("Find pow with initializer variable, which is also scoped to else")
	fmt.Println(
		powV2(3, 2, 10),
		powV2(3, 3, 20),
	)
}
func main() {
	basicIf()
	fmt.Println()
	ifWithInitializer()
	fmt.Println()
	ifElse()
}
