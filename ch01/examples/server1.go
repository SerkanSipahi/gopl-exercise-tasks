package main

import(
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Print("b1\n")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Print("b2\n")

}