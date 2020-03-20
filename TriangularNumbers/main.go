package main

import (
	"fmt"
)

func main() {
	fmt.Println("Highly Divisible Triangular Number")
	var fs []int
	var t int
	for x := 0; x < 1000000; x++ {
		t = triangular(x)
		if t%2 == 0 {
			fs = factors(t)
		} else {
			fs = []int{}
		}
		if x%100 == 0 {
			fmt.Print(x/100, " ")
		}
		if x%1000 == 0 {
			fmt.Println()
			fmt.Println(x / 1000)
		}
		// fmt.Printf("%10d: %10d - %3d\n", x, t, len(fs))
		/*
			for _, n := range fs {
				fmt.Printf("%d, ", n)
			}
			fmt.Println()
		*/
		if len(fs) > 500 {
			break
		}
	}
	fmt.Println()
	fmt.Printf("%10d - %3d\n", t, len(fs))
}

func triangular(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return (n * (n + 1)) / 2
}

func factors(n int) (f []int) {
	for x := 1; x <= n/2; x++ {
		if n%x == 0 {
			f = append(f, x)
		}
	}
	f = append(f, n)
	return
}
