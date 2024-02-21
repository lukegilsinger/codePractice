package main

import (
	"fmt"
)

func main() {
	ifs()
	switches()
}

func ifs() {
	if i := 5; i < 5 { // initializer optional
		fmt.Println("i is less than 5")
	} else if i < 10 {
		fmt.Println("i is less than 10")
	} else {
		fmt.Println("i is greater than or equal to 10")
	}
}

func switches() {
	// i := 2
	switch j := 99; j {
	case 1:
		fmt.Println("first")
	case 2:
		fmt.Println("second")
	default:
		fmt.Println("none")
	}
}
