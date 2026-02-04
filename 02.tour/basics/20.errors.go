package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s at %v", e.What, e.When)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func convert(s string) {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}

	fmt.Println("Converted integer:", i)

}

func catchError() {
	fmt.Println("Functions often return an error value, caller check error to nil.")

	convert("42")
	convert("asdf")
}

func customError() {
	fmt.Println("same returning error pattern can be used for custom errors")

	if err := run(); err != nil {
		fmt.Println(err)
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0

	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		// fmt.Printf("iteration %d: z = %g\n", i+1, z)
	}

	return z, nil
}

func printSqrt(x float64) {
	if r, e := Sqrt(x); e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("Sqrt = %v\n", r)
	}
}

func exercise() {
	fmt.Println("Exercise with return error")
	printSqrt(2)
	printSqrt(-2)
}

func main() {
	catchError()
	fmt.Println()
	customError()
	fmt.Println()
	exercise()
}
