package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("comma", comma("Hello Worlds Fool", 2))
	fmt.Println(isAnAnagram("Astronomer", "Moon starer"))
}

func comma(s string, sep int) string {
	if sep < 1 {
		return "Please enter a positive number"
	}

	var b bytes.Buffer
	n := len(s)
	if n <= sep {
		return s
	}

	reveredStr := reverse(s)

	for idx := 0; idx < len(reveredStr); idx += sep {
		endIdx := idx + sep
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

func isAnAnagram(str1 string, str2 string) bool {
	str1 = strings.ToLower(str1)
	str1 = strings.Replace(str1, " ", "", -1)
	slice1 := strings.Split(str1, "")
	letterMap := map[string]int{}

	for _, letter := range slice1 {
		if _, ok := letterMap[letter]; ok {
			letterMap[letter]++
		} else {
			letterMap[letter] = 1
		}
	}

	str2 = strings.ToLower(str2)
	str2 = strings.Replace(str2, " ", "", -1)
	slice2 := strings.Split(str2, "")

	for _, letter := range slice2 {
		val, ok := letterMap[letter]
		if val == 0 || ok == false {
			return false
		} else {
			letterMap[letter]--
		}
	}

	for key := range letterMap {
		if letterMap[key] != 0 {
			return false
		}
	}

	return true
}
