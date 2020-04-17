package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"io"
)

func main() {
	ex2()
}

func ex1() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetchEx1(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchEx1(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Oooops")
		return
	}
	writeErr := ioutil.WriteFile("otherFileName", str, 0644)
	resp.Body.Close() // don't leak resources
	if err != nil || writeErr != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %s", secs, url)
}

func ex2() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchEx2(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchEx2(url string, ch chan<- string) {
	start := time.Now()
	abort := make(chan struct{})
	tick := time.Tick(5 * time.Second)
	go func() {
		<-tick
		ch <- "Connection Timed Out"
		return
	}()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// issue is that
