package main

import "fmt"

func stacked() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func basic() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

func main() {
	basic()
	fmt.Println()
	stacked()
}
