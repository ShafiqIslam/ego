package main

import "fmt"

func assert() {
	fmt.Println("test whether an interface value holds a specific type")

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// will panic
	// f = i.(float64) // panic
	// fmt.Println(f)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func switches() {
	fmt.Println("several type assertions in series")

	do(21)
	do("hello")
	do(true)
}

func main() {
	assert()
	fmt.Println()
	switches()
}
