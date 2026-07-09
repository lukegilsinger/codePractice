package main

// go run .
// go run main.go

import (
	"fmt"
)

func fun() int {
	println("in f")
	return 0
}

func main() {
	fmt.Println("hello, gophers")

	var myName string
	myName = "Mark"

	var myName2 = "Mike" // inferred type

	myName3 := "Mike" // short declaration syntax

	fmt.Println(myName)
	fmt.Println(myName2)
	fmt.Println(myName3)

	var c = true
	fmt.Println(c)

	d := 3.14
	fmt.Println(d)

	var e int = int(d)
	fmt.Println(e)

	a, b := 10, 5
	var f = a + b
	println(f)

	g := a == b // or !=, <, >= (only work for numbers)
	println(g)

	c = false
	println(c)

	const h = 42
	// h = 43

	const (
		i = "foo"
		j // gets previous value
	)
	println(j)

	const k = 2 * 5
	println(k)

	const l = iota // iota works in constant gropu

	const (
		m = iota
		n
		o = 3 * iota
		p
	)
	println(n)
	println(o)

	fun()
	q := fun()
	println(q)

	r := 42
	s := &r          // is a pointer
	println("s:", s) // address value
	*s = 13          // updates r too. dereferencing
	println("s:", *s)
	println("r:", r)

	t := new(string)
	println("t:", t)
	println("t:", *t)

	u := new(int)
	println(*u)
}
