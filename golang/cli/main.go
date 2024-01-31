// go run . -level INFO

package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"fmt"
	"flag"
)

func main() {

	path := flag.String("path", "myapp.log", "Path to log file")
	level := flag.String("level", "ERROR", "level of error")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil { // errors are pointer values
		log.Fatal(err)

	}
	defer f.Close() // after function closes
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}