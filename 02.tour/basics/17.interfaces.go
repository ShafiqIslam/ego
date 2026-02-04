package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func interfaces() {
	fmt.Println("a set of method signatures")

	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())

	a = &v // a *Vertex implements Abser
	fmt.Println(a.Abs())

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v
	// fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	print()
}

type I2 interface {
	printV2()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) print() {
	fmt.Println(t.S)
}

func (t *T) printV2() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func implicitBehaviour() {
	fmt.Println("A type implements an interface by just implementing its methods")

	var i I = T{"hello"}
	i.print()
}

type F float64

func (f F) print() {
	fmt.Println(f)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interfaceValues() {
	fmt.Println("interface values are a tuple of a value and a concrete type")

	var i I

	i = &T{"Hello"}
	describe(i)
	i.print()

	i = F(math.Pi)
	describe(i)
	i.print()
}

func nilInterface() {
	fmt.Println("If concrete value (not the type) of interface is nil, method will be called with a nil receiver")

	var i I2

	var t *T
	i = t
	describe(i)
	i.printV2()

	i = &T{"hello"}
	describe(i)
	i.printV2()

	var nilI I
	describe(nilI)
	// will panic
	// nilI.print()
}

func emptyInterface() {
	fmt.Println("empty interface may hold values of any type")

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func main() {
	interfaces()
	fmt.Println()
	implicitBehaviour()
	fmt.Println()
	interfaceValues()
	fmt.Println()
	nilInterface()
	fmt.Println()
	emptyInterface()
}
