// Modify the echo program totals to print os.Args[0], the name of the command that invoked it.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("Ex 1")
	ex1()
	fmt.Println("Ex 2")
	ex2()
	fmt.Println("Ex 3")
	ex3()
}

func ex1() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func ex2() {
	s, sep := "", ""
	for idx, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println("idx is", idx, "arg is", arg)
	}
	fmt.Println(s)
}

func ex3() {
	ex3Efficient()
	ex3Original()
}

func ex3Efficient() {
	start := time.Now()

	s := strings.Join(os.Args, " ")

	fmt.Println(s)

	fmt.Println("Time elasped Efficient", time.Since(start).Seconds())
}

func ex3Original() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("Time elasped Original", time.Since(start).Seconds())
}
