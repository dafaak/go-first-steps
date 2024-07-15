package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var ( // asi se puede agrupar
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//0 for numeric types,
//false for the boolean type, and
//"" (the empty string) for strings.

var i int
var f float64
var boolean bool
var s string

// Unlike in C, in Go assignment between items of different type requires an explicit conversion.
var x, y int = 3, 4
var f1 float64 = math.Sqrt(float64(x*x + y*y))
var z1 uint = uint(f1)

// When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
var i1 int
var j = i // j is an int

var i2 = 42          // int
var f2 = 3.142       // float64
var g = 0.867 + 0.5i // complex128

// Constants are declared like variables, but with the const keyword.
// Constants can be character, string, boolean, or numeric values.
// Constants cannot be declared using the := syntax

const Pi = 3.14

func main() {
	fmt.Println("Hola, Go!")
	a, b := swap("hello", "world")

	fmt.Println(a, b)

	fmt.Println(split(10))

	sum := func(a, b int) int { return a + b }(4, 3)

	fmt.Println(sum)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Printf("%v %v %v %q\n", i, f, boolean, s)

	fmt.Println(x, y, z1)

	fmt.Printf("v is of type %T\n", i1)
	fmt.Printf("v is of type %T\n", i2)
	fmt.Printf("v is of type %T\n", j)
	fmt.Printf("v is of type %T\n", f2)
	fmt.Printf("v is of type %T\n", g)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
}

func swap(x, y string) (string, string) {

	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
