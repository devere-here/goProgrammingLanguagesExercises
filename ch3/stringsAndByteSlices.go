package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("comma", comma("Hello Worlds Fool"))
}

func comma(s string) string {
	var b bytes.Buffer
	n := len(s)
	if n <= 3 {
		return s
	}

	reveredStr := reverse(s)

	for idx := 0; idx < len(reveredStr); idx += 3 {
		endIdx := idx + 3
		if endIdx >= len(reveredStr) {
			endIdx = len(reveredStr)
		}

		b.Write([]byte(reveredStr[idx:endIdx] + ","))
	}

	return reverse(strings.TrimRight(b.String(), ","))
}

func reverse(s string) string {
	slice := strings.Split(s, "")
	length := len(slice)

	for idx := 0; idx < length-1-idx; idx++ {
		otherIdx := length - 1 - idx
		slice[idx], slice[otherIdx] = slice[otherIdx], slice[idx]
	}

	return strings.Join(slice, "")
}
