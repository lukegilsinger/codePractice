package main

import (
	"fmt"
	"slices"
	// "slices"
)

func main() {
	// types()
	// arrayDemo()
	// slicesDemo()
	mapDemo()
	fmt.Println("Done")
}

func arrayDemo() {
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"coff", "esp", "capp"}

	fmt.Println(arr)

	fmt.Println(arr)

	arr2 := arr

	arr[1] = "tea"

	fmt.Println(arr, arr2)
}

func slicesDemo() {
	// not fixed size
	// points to an array somewhere else
	// reference data types
	var s []int
	fmt.Println(s) // (nil)
	s = []int{1, 2, 3}

	s[1] = 99
	fmt.Println(s)

	s = append(s, 5, 10, 15)
	fmt.Println(s)

	slices.Delete(s, 1, 3) //remove indeces 1 and 2
	fmt.Println("delete", s)

	s2 := s //pointing to same data
	s2[2] = 99

}

func mapDemo() {
	// grows
	// provide val and key type
	// maps are copied by reference.  point to same data as original.  can use a copy function

	// maps are not comparable m == m2

	var m map[string]int
	fmt.Println(m)

	m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m)

	fmt.Println(m["foo"])

	m["bar"] = 99

	delete(m, "foo")

	m["baz"] = 23

	fmt.Println(m)

	fmt.Println(m["foo"]) // always return resilt.  this will retuen 0

	v, ok := m["foo"] // ok will be true if found.  false if not
	fmt.Println(v, ok)
}

func types() {
	fmt.Println("hi. Doing types")

	var arr [3]int
	fmt.Println(arr)

	arr = [3]int{1, 2, 3}
	fmt.Println(arr[1])

	arr[1] = 99
	fmt.Println(arr[1])

	arr2 := arr
	fmt.Println(arr2)

	arr[0] = 7
	fmt.Println(arr)
	fmt.Println(arr2) //Go copies. not reference

	fmt.Println(arr == arr2)
}
