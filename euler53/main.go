package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Combinatoric Selections")
	fmt.Println("Choose 23 of 10 is", chooseRofN(23, 10))
	million := big.NewInt(1000000)
	c := [][2]int{}
	for n := 1; n <= 100; n++ {
		for r := 0; r <= n; r++ {
			if chooseRofN(n, r).Cmp(million) >= 1 {
				c = append(c, [2]int{n, r})
			}
		}
	}
	/*
		for x := range c {
			fmt.Print(c[x], " ")
			if x%15 == 0 {
				fmt.Println()
			}
		}
	*/
	fmt.Println(len(c))
}
func chooseRofN(n, r int) *big.Int {
	nf := big.NewInt(1)
	rf := big.NewInt(1)
	nLessr := big.NewInt(1)
	inc := big.NewInt(1)
	one := big.NewInt(1)
	for x := 1; x <= n; x++ {
		nf = nf.Mul(nf, inc)
		if x <= r {
			rf = rf.Mul(rf, inc)
		}
		if x <= n-r {
			nLessr = nLessr.Mul(nLessr, inc)
		}
		inc = inc.Add(inc, one)
	}
	// fmt.Println(nf, rf, nLessr)
	return nf.Div(nf, rf.Mul(rf, nLessr))
}
