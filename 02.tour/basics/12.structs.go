package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func literals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		pt = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, pt, v2, v3)
}

func pointer() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

func fields() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func main() {
	fmt.Println(Vertex{1, 2})

	fields()
	pointer()
	literals()
}
