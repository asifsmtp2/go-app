package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Printf("--APPLICATION STARTED--")

	http.HandleFunc("/ping", index)

	http.HandleFunc("/", index)
	http.HandleFunc("/hi", index)

	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("--REQUEST CALLED--")
	io.WriteString(w, "hello from a docker container")
}
