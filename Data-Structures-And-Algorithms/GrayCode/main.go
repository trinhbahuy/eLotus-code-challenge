package main

import (
	"fmt"
	"os"
	"strconv"
)

// grayCode return n-bit gray code sequence
func grayCode(n int) (gc []int) {
	N := 1 << n

	for i := 0; i < N; i++ {
		gc = append(gc, i^(i>>1))
	}

	return
}

func main() {
	// get argument and convert to int
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	// constraints
	if n < 1 || n > 16 {
		panic("n should be greater than 1 and less than 16")
	}

	fmt.Println(grayCode(n))
}
