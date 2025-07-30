package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("what to scream?")

	in := bufio.NewReader(os.Stdin)

	s, _ := in.ReadString('\n') // _ is err that we dont need to use
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	fmt.Println(s + "!")
}
