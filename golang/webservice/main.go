package main
import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //root path. handle all requests
		w.Write([]byte("Hello World"))
	}) 

	err := http.ListenAndServe(":3000", nil) // use default handler

	if err != nil {
		log.Fatal(err)
	}
}
