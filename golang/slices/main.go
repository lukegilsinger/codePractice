package main

import (
	"fmt"
	"slices"
	// "slices"
)

func main() {
	type m struct {
		name string
		age  int
	}

	s := []m{
		{
			name: "one",
			age:  1,
		},
		{
			name: "two",
			age:  2,
		},
	}

	delFunc := func(i m) bool {
		if i.age == 1 {
			return true
		}
		return false
	}
	fmt.Println(s)

	slices.DeleteFunc(s, delFunc)

	fmt.Println(s)
	fmt.Println(len(s))

	s2 := make([]m, 2, 2)
	s2 = append(s2, m{
		name: "three",
		age:  3,
	})

	fmt.Println(s2)
	fmt.Println(len(s2))

	s3 := []m{
		{
			name: "one",
			age:  1,
		},
		{
			name: "two",
			age:  2,
		},
	}
	fmt.Println(s3)
	s3 = slices.Delete(s3, 0, 1)
	fmt.Println(s3)

}
