package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sideProjects/goProgrammingLanguagesExercises/ch2/lengthconv"
	"github.com/sideProjects/goProgrammingLanguagesExercises/ch2/tempconv"
	"github.com/sideProjects/goProgrammingLanguagesExercises/ch2/weightconv"
)

func main() {
	if len(os.Args[1:]) > 0 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			makeConversions(t)
		}
	} else {
		fmt.Println("No number was presented as an argument")
		fmt.Println("Please enter a number")

		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("-> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSuffix(input, "\n")
			fmt.Println("input is", input)
			num, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Err is", err)
			}

			makeConversions(num)
		}
	}
}

func makeConversions(num float64) {
	f := tempconv.Fahrenheit(num)
	c := tempconv.Celsius(num)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	kg := weightconv.Kilogram(num)
	lb := weightconv.Pound(num)
	fmt.Printf("%s = %s, %s = %s\n", kg, weightconv.KgToLb(kg), lb, weightconv.LbToKg(lb))

	m := lengthconv.Meter(num)
	ft := lengthconv.Foot(num)
	fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MToFt(m), ft, lengthconv.FtToM(ft))
}
