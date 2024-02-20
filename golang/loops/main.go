package main

// all for loops in go

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	i := 99

	for { // infinite loop
		fmt.Println(i)
		i += 1
		if i >= 105 {
			break
		}
	}

	i = 5
	for i < 10 { // loops with condition
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ { // counter based.  use initializer statement
		fmt.Println(i)
	}

	// also can iterate over a collection
	menu()
}

func menu() {
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

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, cost := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, cost) // 2 formatting verbs
		}
	}

}
