package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/pic"
)

func basicArrays() {
	fmt.Println("Arrays cannot be resized")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

func slicePointer() {
	fmt.Println("slice does not store any data, it just describes a section of an underlying array.")

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceOfArray() {
	fmt.Println("A slice, on the other hand, is dynamically-sized")

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

func sliceLiterals() {
	fmt.Println("Literal creates the an array, then builds a slice that references it")

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func boundDefaults() {
	fmt.Println("high or low bounds can be omitted")

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func lengthAndCapacity() {
	fmt.Println("capacity of a slice is the size of the underlying array")

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice("a", s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice("b", s)

	// Extend its length.
	s = s[:4]
	printSlice("c", s)

	// Drop its first two values.
	s = s[2:]
	printSlice("d", s)
}

func printSlice(n string, s []int) {
	fmt.Printf("%s, len=%d cap=%d %v\n", n, len(s), cap(s), s)
}

func nilSlices() {
	fmt.Println("The zero value of a slice is nil")

	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func makingSlices() {
	fmt.Println("built-in make function to create dynamically-sized arrays")

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func sliceOfSlices() {
	fmt.Println("Slices can contain any type, including other slices")

	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendingSlices() {
	fmt.Println("append new elements to a slice")

	var s []int
	printSlice("a", s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice("b", s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice("c", s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice("d", s)
}

func ranges() {
	fmt.Println("iterates over a slice or map, wit or without index")

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d, ", i, v)
	}
	fmt.Println()

	for _, value := range pow {
		fmt.Printf("%d, ", value)
	}
	fmt.Println()
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8(x ^ y)
		}
		pic[y] = row
	}

	return pic
}

func exerciseSlice() {
	fmt.Println("An exercise of generating an image as a 2D slice")
	pic.Show(Pic)
}

func main() {
	basicArrays()
	fmt.Println()
	sliceOfArray()
	fmt.Println()
	slicePointer()
	fmt.Println()
	sliceLiterals()
	fmt.Println()
	boundDefaults()
	fmt.Println()
	lengthAndCapacity()
	fmt.Println()
	nilSlices()
	fmt.Println()
	makingSlices()
	fmt.Println()
	sliceOfSlices()
	fmt.Println()
	appendingSlices()
	fmt.Println()
	ranges()
	fmt.Println()
	exerciseSlice()
}
