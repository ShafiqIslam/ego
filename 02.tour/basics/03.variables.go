package main

import "fmt"

var c, python, java bool

func initializer() {
	fmt.Println("one per variable initializer; which can omit type declaration")

	var i, j int = 1, 2
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func shortDeclaration() {
	fmt.Println(":= can only be at function level")

	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

func statement() {
	fmt.Println("var statement can be at package or function level")

	var i int
	fmt.Println(i, c, python, java)
}

func main() {
	statement()
	fmt.Println()
	initializer()
	fmt.Println()
	shortDeclaration()
}
