package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addV2(x, y int) int {
	return x + y
}

func parameters() {
	fmt.Println("function can have omitted type in case of consecutive similar params")

	fmt.Println(add(42, 13))
	fmt.Println(addV2(42, 23))
}

func swap(x, y string) (string, string) {
	return y, x
}

func multiReturn() {
	fmt.Println("functions can return multiple values")

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func namedReturn() {
	fmt.Println("naked return statement returns the named return values")
	fmt.Println(split(17))
}

func main() {
	parameters()
	fmt.Println()
	multiReturn()
	fmt.Println()
	namedReturn()
}
