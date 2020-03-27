package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Powerful Digit Sum")
	a := big.NewInt(1)
	b := big.NewInt(1)
	c := new(big.Int)
	one := big.NewInt(1)
	fmt.Println(a.Exp(a, b, nil).String())
	fmt.Println("Sum:", sum(a.String()))
	max := 0
	for x := 0; x < 99; x++ {
		b = big.NewInt(1)
		for y := 0; y < 99; y++ {
			s := sum(c.Exp(a, b, nil).String())
			if max < s {
				max = s
			}
			b.Add(b, one)
		}
		a.Add(a, one)
	}
	fmt.Println("Max digits sum:", max, a, b)
}

func sum(s string) int {
	sum := 0
	for _, c := range s {
		sum += int(c) - int('0')
	}
	return sum
}
