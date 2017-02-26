package main

import (
	"fmt"
	"net/http"
	"github.com/labstack/gommon/log"
)

func handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "HEADER[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "HOST = %q\n", r.Host)
	fmt.Fprintf(w, "RemoveAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
