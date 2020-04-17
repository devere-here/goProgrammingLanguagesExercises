package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if (strings.HasPrefix(url, prefix) == false) {
			url = prefix + url
		}

		fmt.Println("the url is", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		size, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("size is", size, "in bytes")
		fmt.Println("response code is", resp.Status)
	}
}