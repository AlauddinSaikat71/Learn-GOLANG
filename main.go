package main

import (
	"fmt"
	"reflect"
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

// constant example
func constants() {
	const myConstant1 int = 42
	fmt.Printf("%v %T\n", myConstant1, myConstant1)

	// type not necessary
	const myConstant2 = 15
	fmt.Printf("%v %T\n", myConstant2, myConstant2)

	var b int16 = 2
	fmt.Printf("%v  %T\n", myConstant2+b, myConstant2+b)

	//iota
	const (
		j = iota
		k
		l
	)
	fmt.Printf("j=%v  k=%v  l=%v  \n", j, k, l)

	// another const block set value from 0 again
	const (
		j2 = iota
	)
	fmt.Printf("j2=%v   \n", j2)

	//const block for kb,mb,gb,tb
	const (
		_  = iota //ignore first value by assigning to blank identifier
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
	)
	fileSize := 4000000000.
	fmt.Printf("%.2f GB\n", fileSize/GB)

	//by bitshifting checking roles.But i didn't understand how i happened
	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinancials

		canSeeAfrica
		canSeeAsia
		canSeeEurope
	)
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin ? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ ? %v\n", isHeadquarters&roles == isHeadquarters)
}

func arrays() {

	grades := [3]int{10, 20, 23}
	//another way to declare an array,dots intead of size
	//grades := [...]int{10, 20, 23}
	fmt.Printf("Grades : %v\n", grades)

	var students [3]string
	fmt.Printf("Studnets : %v\n", students)
	students[0] = "Saikat"
	fmt.Printf("Studnets : %v\n", students)
	students[1] = "Lisa"
	fmt.Printf("Studnets #1 : %v\n", students[1])

	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 2, 3}
	identityMatrix[1] = [3]int{1, 0, 5}
	identityMatrix[2] = [3]int{1, 4, 9}
	fmt.Println(identityMatrix)

	// declare array b and assign a
	a := [3]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	//declare array c and assign address of a
	c := &a
	c[1] = 5 // c change & also a change
	fmt.Println(a)
	fmt.Println(c)

}

func slices() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	fmt.Printf("Length : %v\n", len(a))
	fmt.Printf("Length : %v\n", cap(a))

	b := a[:]   //slice of all elements
	c := a[3:]  // slice from 4th ele to end
	d := a[:6]  // slice of first 6 ele
	e := a[3:6] // slice of 4th,5th,6th ele
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// change one ele from a and change on all slices
	fmt.Println(" after changing from slice 'a' ")
	a[5] = 42
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

}

func maps() {
	family := make(map[string]int)
	family = map[string]int{
		"siakat": 10,
		"rakib":  13,
		"abir":   20,
	}
	fmt.Println(family)

	// after adding mili
	family["Mili"] = 30
	fmt.Println(family)

	delete(family, "Mili")
	fmt.Println(family)

	pop := family["Shawon"]
	fmt.Println(pop)

	pop, ok := family["abir"]
	fmt.Println(pop, ok)
}

func structures() {
	type Doctor struct {
		number        int
		actorName     string
		companiesName []string
	}

	aDoctor := Doctor{
		number:    3,
		actorName: "Saikat",
		companiesName: []string{
			"rakib",
			"abir",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number)

	bDoctor := Doctor{
		10,
		"rafi",
		[]string{
			"muslim high",
		},
	}
	fmt.Println(bDoctor)

	//another way to declare struct
	cDoctor := struct {
		name string
		age  int
	}{name: "John", age: 30}
	fmt.Println(cDoctor)

	//go doesn't support inheritance
	// but it has composition
	type Animal struct {
		Name   string
		Origin string
	}
	type Bird struct {
		Animal
		SpeedKPH float32
		CanFly   bool
	}
	aBird := Bird{}
	aBird.Name = "Emu"
	aBird.Origin = "Australia"
	aBird.SpeedKPH = 48
	aBird.CanFly = true
	fmt.Println(aBird)

	// another bird
	bBird := Bird{
		Animal:   Animal{Name: "emu", Origin: "austra"},
		SpeedKPH: 30,
		CanFly:   false,
	}
	fmt.Println(bBird)

	//tags
	type Dog struct {
		Name   string `required max:"100"`
		Origin string
	}
	t := reflect.TypeOf(Dog{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}

func main() {

	var mode int
	fmt.Println("Enter your mode : ")
	fmt.Scan(&mode)
	fmt.Println(mode)

	switch mode {
	case 1:
		variable_primitives()
	case 2:
		constants()
	case 3:
		arrays()
	case 4:
		slices()
	case 5:
		maps()
	case 6:
		structures()
	}

	//
	//
	//
	//
}
