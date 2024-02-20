package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("LEARNING")

	// maps()
	// structs()
	demo()

}

func demo() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice) // we don't know what to do with this yet!

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.65, "medium": 1.80, "large": 1.95}},
		{name: "Espresso", prices: map[string]float64{"single": 1.90, "double": 2.25, "triple": 2.55}},
	}

	fmt.Println(menu)
}

func maps() {
	var m = map[string][]string{
		"coffee": {"coof", "esp", "cap"},
		"tea":    {"hot", "chai"},
	}
	fmt.Println(m)

	m["other"] = []string{"choco", "soda"}
	fmt.Println(m)

	delete(m, "tea")
	fmt.Println(m)

	m2 := m // reference same data
	fmt.Println(m2)
}

func structs() {
	var s struct {
		name string
		id   int
	}
	fmt.Println(s)

	s.name = "arthur"
	fmt.Println(s)

	// can create custom type based on struct
	type myStruct struct {
		name string
		id   int
	}
	// fmt.Println(myStruct)

	var s2 myStruct
	fmt.Println(s2)

	s2 = myStruct{
		name: "art",
		id:   42,
	}
	fmt.Println(s2)

	s3 := s2 //value types.  like arrays
	s2.name = "Tri"
	fmt.Println(s2, s3)
}
