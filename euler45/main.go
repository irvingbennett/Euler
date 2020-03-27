package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Triangular, Pentagonal & Hexagonal")
	for x := 286; x < 99999; x++ {
		n := triangle(x)
		p := isPentagonal(n)
		h := isHexagonal(n)
		if p > 0 && h > 0 {
			// found the number
			fmt.Printf("Triangular of %d is %d\n", x, n)
			fmt.Println(n, "is algo pentagonal", p)
			fmt.Println(n, "is also hexagonal", h)
			break
		}
	}

}

func triangle(n int) int {
	return n * (n + 1) / 2
}

func isPentagonal(x int) int {
	var n float64
	n = (math.Sqrt(float64(24*x+1)) + 1) / 6
	// fmt.Println(x, n, n-float64(int(n)))

	if n-float64(int(n)) == 0 {
		return int(n)
	}
	return -1
}

func isHexagonal(x int) int {
	var n float64
	n = ((math.Sqrt(float64(8*x + 1))) + 1) / 4

	if n-float64(int(n)) == 0 {
		return int(n)
	}
	return -1

}
