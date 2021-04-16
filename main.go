package main

import (
	"fmt"
	"strconv"
)

func variable_primitives() {
	var i int = 42
	fmt.Printf(" i = %v  %T\n", i, i)

	var k float32 = 42.5
	fmt.Printf(" k = %v  %T\n", k, k)

	l := 3
	fmt.Printf(" l = %v  %T\n", l, l)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf(" j = %v , %T\n", j, j)

	var n bool = true
	fmt.Printf(" n = %v  %T\n", n, n)

	var m = 1 == 2 // 1==2 means false
	fmt.Printf(" m = %v  %T\n", m, m)
	m = 1 == 1 //1==1 means true
	fmt.Printf(" m = %v  %T\n", m, m)

	a := 10 //1010
	b := 3  //0011
	//math operations
	println(a + b)
	println(a - b)

	println(a & b) // binary and
	println(a | b) // binary or
	println(a ^ b)
	println(a &^ b) // binary and

	a = 16
	fmt.Println(a << 2)
	fmt.Println(a >> 2)

	//complex number
	c := 1 + 2i
	d := 2 + 5.2i
	fmt.Println(c + d)

	var complexnumber complex64 = 1 + 2i
	fmt.Printf("%v  %T\n", real(complexnumber), real(complexnumber))
	fmt.Printf("%v  %T\n", imag(complexnumber), imag(complexnumber))

	var complexnumber2 complex128 = complex(5, 12)
	fmt.Printf("%v  %T\n", complexnumber2, complexnumber2)

	//rune
	var r rune = 'a'
	fmt.Printf("%v  %T\n", r, r)

}

func constants() {
	const myConstant1 int = 42
	fmt.Printf("%v %T\n", myConstant1, myConstant1)
}

func main() {

	//variable_primitives()
	constants()
}
