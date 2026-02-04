package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

func nilMaps() {
	fmt.Println("nil map has no keys, nor can keys be added")

	var m map[string]int
	fmt.Println(m["a"])
	v, ok := m["a"]
	fmt.Println(v, ok)

	// will panic
	// m["a"] = 1
	// fmt.Println(m["a"])
}

func makeMap() {
	fmt.Println("make returns a map of given type, initialized and ready for use")
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

func literals() {
	fmt.Println("Map literals are like struct literals, but the keys are required")
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	var m2 = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m2)
}

func mutation() {
	fmt.Println("mutating maps")

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("Insert; value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("Update; value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("Delete; value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("Check presense; value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	words := strings.Fields(s)
	for _, w := range words {
		m[w]++
	}

	return m
}

func mapExercise() {
	fmt.Println("an exercise of frequency maps")
	wc.Test(WordCount)
}

func main() {
	nilMaps()
	fmt.Println()
	makeMap()
	fmt.Println()
	literals()
	fmt.Println()
	mutation()
	fmt.Println()
	mapExercise()
}
