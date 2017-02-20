package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

/**
 * The default http.Get has no timeout. So if the server does not respond your program will wait for ever!
 * see solution: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779#.zfgnbt2eu
 */
var netClient = &http.Client{
	Timeout: time.Second * 2,
}

func main(){

	args := os.Args[0:]

	start := time.Now()
	ch := make(chan string)

	for _, url := range args {
		go fetch(url, ch)
	}

	for range args {
		fmt.Println(<-ch)
	}

	fmt.Printf("%2fs elapsed\n", time.Since(start).Seconds())

}

func fetch(url string, ch chan<- string){

	start:= time.Now()
	resp, err := netClient.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d %s", secs, nbytes, url)
}