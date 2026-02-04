package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("iteration %d: z = %g\n", i+1, z)
	}

	return z
}

func SqrtV2(x float64) float64 {
	i, z := 1, 1.0

	for {
		prev := z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("iteration %d: z = %g\n", i, z)
		i += 1
		if math.Abs(z-prev) < 1e-5 {
			break
		}
	}

	return z
}

func main() {
	fmt.Println("With fixed 10 iteration:")
	fmt.Println(Sqrt(2))
	fmt.Println("\nWith non-fixed iteration & proximity check:")
	fmt.Println(SqrtV2(100))
}
