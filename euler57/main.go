package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {
	fmt.Println("Square Root Convergents")
	one := big.NewRat(1, 1)
	j := big.NewRat(1, 1)
	count := 0
	for x := 0; x < 1000; x++ {
		j.Add(one, expand(x))
		parts := strings.Split(j.String(), "/")
		if len(parts[0]) > len(parts[1]) {
			count++
		}
		if x < 20 {
			fmt.Println(j.FloatString(10), j, parts[0], parts[1])
		}

	}
	fmt.Println("Count is:", count)
}

func expand(n int) *big.Rat {
	j := big.NewRat(1, 1)
	two := big.NewRat(2, 1)
	if n == 0 {
		return j.Quo(j, big.NewRat(2, 1))
	}
	return j.Quo(j, two.Add(two, expand(n-1)))

}
