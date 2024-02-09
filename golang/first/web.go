package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	println("web")
	http.HandleFunc("/", Handler)
	http.ListenAndServe("localhost:3000", nil) // go will prvide front controller
}

func Handler(w http.ResponseWriter, r *http.Request) {
	println("handler")
	f, _ := os.Open("./meun.txt")
	io.Copy(w, f)
}
