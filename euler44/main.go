package main

import (
	"fmt"
	"math"
)

var pns map[int]bool

func main() {

	pns = make(map[int]bool, 0)
	candidates := [][]int{}
	fmt.Println("Pentagon Numbers")
	fmt.Println("Pentagon 3 is", pn(3))
	fmt.Println("92 is pentagonal", isPentagonal(92))

	fmt.Println("8 is pentagonal", isPentagonal(8))

	for i, c := 1, 1; c <= 25000000; i, c = i+3, c+i+3 {
		pns[c] = true
	}
	for j := 1; j < 5000; j++ {
		Pj := j * (3*j - 1) / 2
		for k := j + 1; k < 5000; k++ {
			Pk := k * (3*k - 1) / 2
			pn := check(Pj, Pk)
			if len(pn) > 0 {
				candidates = append(candidates, pn)
			}
		}

	}

	for x := range candidates {
		fmt.Println(candidates[x])
		fmt.Println(candidates[x][1] - candidates[x][0])
		if x > 100 {
			break
		}
	}

}

// returns pentagon number n
func pn(n int) int {
	return n * (3*n - 1) / 2
}

func check(p1, p2 int) []int {
	var plus, minus bool
	if pns[p1+p2] {
		//fmt.Printf("Is plus: %d %d \n", p1, p2)
		plus = true
	}
	if pns[p2-p1] {
		//fmt.Printf("Is minus: %d %d \n", p1, p2)
		minus = true
	}

	if plus &&
		minus {
		/*
			fmt.Println("------------------------------------")
			fmt.Printf("Is pentagonal: %d %d \n", p1, p2)
			fmt.Println("------------------------------------")
		*/
		return []int{p1, p2}
	}
	return []int{}

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
