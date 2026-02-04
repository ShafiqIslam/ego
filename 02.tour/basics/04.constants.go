package main

import "fmt"

const Pi = 3.14

func simpleConstants() {
	fmt.Println("Untyped constant takes the type needed by its context")
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func numericConstants() {
	fmt.Println("Numeric constants represent values of arbitrary precision and do not overflow")
	// fmt.Printf("%v", Big)
	// fmt.Println(needInt(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

func main() {
	simpleConstants()
	fmt.Println()
	numericConstants()
}
