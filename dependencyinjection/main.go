package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	greet(w, "World")
}
