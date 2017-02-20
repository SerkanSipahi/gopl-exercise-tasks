package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"io"
)

const (
	prefix = "http://"
)

func main(){

	for _, url := range os.Args[1:] {

		if strings.HasPrefix(url, prefix) == false {
			url = prefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		fmt.Fprintf(os.Stdout, "fetch: response status: %s\n", resp.Status)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s%v\n", url, err)
		}
	}

}