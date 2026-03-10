package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {
// 	s := "👩🏾‍⚕️"                  // example "dusky"
// 	fmt.Println("bytes:", len(s)) // number of bytes
// 	fmt.Println(utf8.RuneCountInString(s))
// }

import (
	"fmt"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

func main() {
	s := "👩🏾‍⚕️" // example "dusky" doctor emoji

	fmt.Println("bytes:", len(s))
	fmt.Println("runes:", utf8.RuneCountInString(s))

	// Split into grapheme clusters
	var clusters []string
	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		clusters = append(clusters, gr.Str())
	}

	fmt.Println("grapheme clusters count:", len(clusters))
	fmt.Printf("clusters: %#v\n", clusters)

	// Treat the first grapheme cluster as a single "character"
	if len(clusters) > 0 {
		first := clusters[0]
		fmt.Println("first cluster (as string):", first)
		fmt.Println("first cluster bytes:", len(first))
		fmt.Println("first cluster runes:", utf8.RuneCountInString(first))
		// conceptually this is the "length 1" character in grapheme terms
	}
}
