package popcount

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}

	fmt.Println("pc is", pc)
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountEx1 returns the population count (number of set bits) of x.
func PopCountEx1(x uint64) int {
	count := 0
	for idx := 0; idx < 8; idx++ {
		count += int(pc[byte(x>>(idx*8))])
	}

	return count
}

func PopCountEx2(x uint64) int {
	count := 0
	for idx := 0; idx < 64; idx++ {
		count += int(pc[idx])
	}

	return count
}
