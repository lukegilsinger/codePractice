package main

import (
	"fmt"
	"io"
	// "log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, world.")

	http.HandleFunc("/", Handler) //root path. handle all requests

	http.ListenAndServe("localhost:3000", nil) // use default handler

	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Goodbye, world.")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling")
	f, _ := os.Open("menu.txt")
	io.Copy(w, f)

}
